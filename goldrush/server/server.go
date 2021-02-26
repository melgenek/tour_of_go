package main

import (
	"../models"
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
				if field.Cells[i][j].HasGold {
					sum += 1
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

func dig(field *models.Field) Endpoint {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			digReq := models.DigRequest{}
			json.NewDecoder(req.Body).Decode(&digReq)

			if field.Dig(digReq.PosX, digReq.PosY) {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]string{"Gold!"})
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]string{})
			}
		}
	}
}

func cash() Endpoint {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]int{1})
		}
	}
}

func main() {
	field := models.CreateField()
	licenseList := models.NewLicenses()

	http.HandleFunc("/explore", explore(field))
	http.HandleFunc("/licenses", licenses(licenseList))
	http.HandleFunc("/dig", dig(field))
	http.HandleFunc("/cash", cash())

	http.ListenAndServe(":8000", nil)
}
