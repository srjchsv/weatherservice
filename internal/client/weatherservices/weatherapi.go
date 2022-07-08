package weatherservices

import (
	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/api"
	"github.com/srjchsv/weatherservice/internal/utils"
)

var WeatherApi WeatherApiResponseStruct

type WeatherApiResponseStruct struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Celsius float64 `json:"temp_c"`
	} `json:"current"`
}

func (ws *WeatherApiResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	url := api.Configs.Url.WeatherApi + "?q=" + url.QueryEscape(location)
	res, err := utils.RequestResponseRapidApi(ctx, url, api.Configs.ApiHost.WeatherApi, api.Configs.RapidApiKey)
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

//STUBS

// func (ws *WeatherApiResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
// 	file, err := ioutil.ReadFile("./stubsJSON/responseWeatherApi")
// 	if err != nil {
// 		return utils.Data{}, err
// 	}
// 	var d WeatherApiResponseStruct

// 	_ = json.Unmarshal(file, &d)

// 	stdData := utils.Data{
// 		Name:        "WeatherApi",
// 		Location:    d.Location.Name,
// 		Temperature: d.Current.Celsius,
// 	}
// 	return stdData, nil
// }
