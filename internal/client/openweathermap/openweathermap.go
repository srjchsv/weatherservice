package openweathermap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	apiConfig, err := utils.LoadApiConfig(ctx, "./configs/.configs")
	if err != nil {
		return utils.Data{}, err
	}

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=" + location + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.Data{}, err
	}
	
	req.Header.Add("X-RapidAPI-Key", apiConfig.OpenWeatherMapApiKey)
	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d OpenWeatherResponseStruct

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
	file, err := ioutil.ReadFile("./stubsJSON/response.json")
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
