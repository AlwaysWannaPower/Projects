package weather

import (
	"App/geo"
	"fmt"
	"net/url"
)

func GetWeather(geo geo.GeoData) string {
	baseUrl, err := url.Parse("https://wttr.in" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", string(format))

}
