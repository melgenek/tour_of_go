package client

import (
	"../models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MineClient struct {
	Host string
}

func (client *MineClient) url(path string) string {
	return fmt.Sprintf("http://%s:8000/%s", client.Host, path)
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

func (client *MineClient) IssueLicense() (models.License, error) {
	req, _ := json.Marshal([]string{})
	license := models.License{}
	err := client.safePost("licenses", req, successfulResponse, func(res *http.Response) error {
		return json.NewDecoder(res.Body).Decode(&license)
	})
	return license, err
}

func (client *MineClient) Cash(gold string) error {
	_, err := http.Post(client.url("cash"), "application/json", bytes.NewBufferString(gold))
	return err
}

type isSuccess func(*http.Response) bool
type callback func(*http.Response) error

func (client *MineClient) safePost(path string, req []byte, isSuccess isSuccess, responseCallback callback) error {
	res, err := http.Post(client.url(path), "application/json", bytes.NewBuffer(req))
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
