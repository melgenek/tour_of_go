package main

import (
	"./models"
	"encoding/json"
	"net/http"
)

type Endpoint = func(w http.ResponseWriter, req *http.Request)

func explore(field *models.Field) Endpoint {
	return func(w http.ResponseWriter, req *http.Request) {
		area := models.Area{}
		json.NewDecoder(req.Body).Decode(&area)

		sum := 0
		for i := area.PosX; i < area.PosX+area.SizeX; i++ {
			for j := area.PosY; j < area.PosY+area.SizeY; j++ {
				for _, n := range field[i][j] {
					if n {
						sum += 1
					}
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ExploreResp{Area: area, Amount: sum})
	}
}

func licenses(licenses *models.Licenses) Endpoint {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			newLicense := licenses.Generate()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newLicense)
		}
	}
}

func main() {
	field := models.CreateField()
	licenseList := models.NewLicenses()

	http.HandleFunc("/explore", explore(&field))
	http.HandleFunc("/licenses", licenses(licenseList))

	http.ListenAndServe(":8000", nil)
}
