package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goldrush/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/expfmt"
)

type MineClient struct {
	host     string
	client   *http.Client
	registry *prometheus.Registry
}

func NewMineClient(host string) *MineClient {
	registry := prometheus.NewRegistry()

	return &MineClient{
		host: host,
		client: &http.Client{
			Transport: instrumentTransport(http.DefaultTransport, registry),
		},
		registry: registry,
	}
}

func instrumentTransport(next http.RoundTripper, registry *prometheus.Registry) promhttp.RoundTripperFunc {
	duration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_histogram_seconds",
			Buckets: []float64{.1, 1},
		},
		[]string{"code", "method", "path"},
	)
	inFlight := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "in_flight_total",
		},
	)

	registry.MustRegister(duration)
	registry.MustRegister(inFlight)

	return func(r *http.Request) (*http.Response, error) {
		inFlight.Inc()
		defer inFlight.Dec()
		start := time.Now()

		resp, err := next.RoundTrip(r)

		labels := prometheus.Labels{}
		labels["method"] = strings.ToLower(r.Method)
		labels["path"] = strings.ToLower(r.URL.Path)
		if err == nil {
			labels["code"] = strconv.Itoa(resp.StatusCode)
		} else {
			labels["code"] = strconv.Itoa(555)
		}
		duration.With(labels).Observe(time.Since(start).Seconds())

		return resp, err
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

func (client *MineClient) Explore(posX int, posY int) (models.ExploreResp, error) {
	req, _ := json.Marshal(models.Area{PosX: posX, PosY: posY, SizeX: 1, SizeY: 1})
	exploreRes := models.ExploreResp{}
	err := client.safePost("explore", req, successfulResponse, func(res *http.Response) error {
		return json.NewDecoder(res.Body).Decode(&exploreRes)
	})
	return exploreRes, err
}

func (client *MineClient) Dig(posX int, posY int, depth int, licenseId int) ([]string, error) {
	req, _ := json.Marshal(models.DigRequest{PosX: posX, PosY: posY, Depth: depth, LicenseID: licenseId})
	var gold []string
	err := client.safePost("dig", req, func(res *http.Response) bool {
		return res.StatusCode == 200 || res.StatusCode == 404
	}, func(res *http.Response) error {
		if res.StatusCode == 200 {
			return json.NewDecoder(res.Body).Decode(&gold)
		} else {
			return nil
		}
	})
	return gold, err
}

func (client *MineClient) IssueLicense(cash []int) (models.License, error) {
	req, _ := json.Marshal(cash)
	license := models.License{}
	err := client.safePost("licenses", req, successfulResponse, func(res *http.Response) error {
		return json.NewDecoder(res.Body).Decode(&license)
	})
	return license, err
}

func (client *MineClient) Cash(gold string) ([]int, error) {
	req, _ := json.Marshal(gold)
	var wallet []int
	err := client.safePost("cash", req, successfulResponse, func(res *http.Response) error {
		return json.NewDecoder(res.Body).Decode(&wallet)
	})
	return wallet, err
}

type isSuccess func(*http.Response) bool
type callback func(*http.Response) error

func (client *MineClient) safePost(path string, req []byte, isSuccess isSuccess, responseCallback callback) error {
	res, err := client.client.Post(client.url(path), "application/json", bytes.NewBuffer(req))
	if err != nil {
		return fmt.Errorf("The http error was '%s'. Path: /%s", err, path)
	} else if !isSuccess(res) {
		return fmt.Errorf("The status was '%d'. Path: /%s", res.StatusCode, path)
	} else {
		callbackErr := responseCallback(res)
		if callbackErr != nil {
			return fmt.Errorf("The parsing error was '%s'. Path: /%s", callbackErr, path)
		} else {
			return nil
		}
	}
}

func successfulResponse(res *http.Response) bool {
	return res.StatusCode == 200
}
