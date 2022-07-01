package weatherservices

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/pkg/utils"
)

type OpenWeatherResponseStruct struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func OpenWeatherMapApi(ctx *gin.Context, location string) (utils.Data, error) {

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=" + url.QueryEscape(location) + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"

	apiHost := "community-open-weather-map.p.rapidapi.com"

	var d OpenWeatherResponseStruct

	res, err := utils.RequestResponseRapidApi(ctx, url, apiHost)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return utils.Data{}, err
	}

	stdData := utils.Data{
		Name:        "OpenWeatherMapApi",
		Location:    d.Name,
		Temperature: d.Main.Celsius,
	}

	return stdData, nil
}

////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============

func OpenWeatherMapApiSTUB(ctx *gin.Context, location string) (utils.Data, error) {

	//save response to file
	file, err := ioutil.ReadFile("./stubsJSON/responseOpenWeatherMap")
	if err != nil {
		return utils.Data{}, err
	}

	var d OpenWeatherResponseStruct

	_ = json.Unmarshal([]byte(file), &d)

	stdData := utils.Data{
		Name:        "OpenWeatherMapApi",
		Location:    d.Name,
		Temperature: d.Main.Celsius,
	}

	return stdData, nil
}
