package geo

import (
	"errors"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func getMyLocation(city string) (*GeoData, error) {
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Not200")
	}
	resp.Body
}
