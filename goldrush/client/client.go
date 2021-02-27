package client

import (
	"bytes"
	"fmt"
	"goldrush/models"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigFastest

type MineClient struct {
	host     string
	registry *prometheus.Registry
	duration *prometheus.HistogramVec
	inFlight prometheus.Gauge
}

func NewMineClient(host string) *MineClient {
	registry := prometheus.NewRegistry()
	duration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_histogram_seconds",
			Buckets: []float64{.1, 1},
		},
		[]string{"code", "path"},
	)
	inFlight := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "in_flight_total",
		},
	)

	registry.MustRegister(duration)
	registry.MustRegister(inFlight)

	return &MineClient{
		host:     host,
		registry: registry,
		duration: duration,
		inFlight: inFlight,
	}
}

func (client *MineClient) PrintMetrics() {
	gathering, err := client.registry.Gather()
	if err != nil {
		fmt.Println(err)
	}

	out := &bytes.Buffer{}
	for _, mf := range gathering {
		mf.Help = nil
		if _, err := expfmt.MetricFamilyToText(out, mf); err != nil {
			panic(err)
		}
	}

	fmt.Print(out.String())
}

func (client *MineClient) url(path string) string {
	return fmt.Sprintf("http://%s:8000/%s", client.host, path)
}

func (client *MineClient) Explore(area *models.Area) (models.ExploreResp, error) {
	req, _ := json.Marshal(area)
	exploreRes := models.ExploreResp{}
	err := client.safePost("explore", req, successfulResponse, func(res *fasthttp.Response) error {
		return json.Unmarshal(res.Body(), &exploreRes)
	})
	return exploreRes, err
}

func (client *MineClient) Dig(posX int, posY int, depth int, licenseId int) ([]string, error) {
	req, _ := json.Marshal(models.DigRequest{PosX: posX, PosY: posY, Depth: depth, LicenseID: licenseId})
	var gold []string
	err := client.safePost("dig", req, func(res *fasthttp.Response) bool {
		return res.StatusCode() == 200 || res.StatusCode() == 404
	}, func(res *fasthttp.Response) error {
		if res.StatusCode() == 200 {
			return json.Unmarshal(res.Body(), &gold)
		} else {
			return nil
		}
	})
	return gold, err
}

func (client *MineClient) IssueLicense(cash []int) (models.License, error) {
	req, _ := json.Marshal(cash)
	license := models.License{}
	err := client.safePost("licenses", req, successfulResponse, func(res *fasthttp.Response) error {
		return json.Unmarshal(res.Body(), &license)
	})
	return license, err
}

func (client *MineClient) Cash(gold string) ([]int, error) {
	req, _ := json.Marshal(gold)
	var wallet []int
	err := client.safePost("cash", req, successfulResponse, func(res *fasthttp.Response) error {
		return json.Unmarshal(res.Body(), &wallet)
	})
	return wallet, err
}

type isSuccess func(*fasthttp.Response) bool
type callback func(*fasthttp.Response) error

var strPost = []byte("POST")
var jsonContentType = []byte("application/json")

func (client *MineClient) safePost(path string, reqBody []byte, isSuccess isSuccess, responseCallback callback) error {
	client.inFlight.Inc()
	defer client.inFlight.Dec()
	start := time.Now()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetBody(reqBody)
	req.Header.SetMethodBytes(strPost)
	req.Header.SetContentTypeBytes(jsonContentType)
	req.SetRequestURI(client.url(path))

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	err := fasthttp.Do(req, res)

	var resultError error
	if err != nil {
		resultError = fmt.Errorf("The http error was '%s'. Path: /%s", err, path)
	} else if !isSuccess(res) {
		resultError = fmt.Errorf("The status was '%d'. Path: /%s", res.StatusCode(), path)
	} else {
		callbackErr := responseCallback(res)
		if callbackErr != nil {
			resultError = fmt.Errorf("The parsing error was '%s'. Path: /%s", callbackErr, path)
		}
	}

	labels := prometheus.Labels{}
	labels["path"] = strings.ToLower(string(req.URI().Path()))
	if err == nil {
		labels["code"] = strconv.Itoa(res.StatusCode())
	} else {
		labels["code"] = strconv.Itoa(555)
	}
	client.duration.With(labels).Observe(time.Since(start).Seconds())

	return resultError
}

func successfulResponse(res *fasthttp.Response) bool {
	return res.StatusCode() == fasthttp.StatusOK
}
