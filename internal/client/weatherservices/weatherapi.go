package weatherservices

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/pkg/utils"
)

type WeatherApiResponseStruct struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Celsius float64 `json:"temp_c"`
	} `json:"current"`
}

func WeatherApi(ctx *gin.Context, location string) (utils.Data, error) {

	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + url.QueryEscape(location)
	apiHost := "weatherapi-com.p.rapidapi.com"

	res, err := utils.RequestResponseRapidApi(ctx, url, apiHost)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d WeatherApiResponseStruct

	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return utils.Data{}, err
	}

	stdData := utils.Data{
		Name:        "WeatherApi",
		Location:    d.Location.Name,
		Temperature: d.Current.Celsius,
	}
	return stdData, nil
}

////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============

func WeatherApiSTUB(ctx *gin.Context, location string) (utils.Data, error) {
	file, err := ioutil.ReadFile("./stubsJSON/responseWeatherApi")
	if err != nil {
		return utils.Data{}, err
	}
	var d WeatherApiResponseStruct

	_ = json.Unmarshal([]byte(file), &d)

	stdData := utils.Data{
		Name:        "WeatherApi",
		Location:    d.Location.Name,
		Temperature: d.Current.Celsius,
	}
	return stdData, nil
}
