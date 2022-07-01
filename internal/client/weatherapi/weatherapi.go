package weatherapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	apiConfig, err := utils.LoadApiConfig(ctx, "./configs/.configs")
	if err != nil {
		return utils.Data{}, err
	}
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.Data{}, err
	}

	req.Header.Add("X-RapidAPI-Key", apiConfig.OpenWeatherMapApiKey)
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
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
