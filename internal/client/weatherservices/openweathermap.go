package weatherservices

import (
	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/api"
	"github.com/srjchsv/weatherservice/internal/utils"
)

var OpenWeatherMapApi OpenWeatherResponseStruct

type OpenWeatherResponseStruct struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func (ws *OpenWeatherResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	url := api.Configs.Url.OpenWeatherMap + "?q=" + url.QueryEscape(location) + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"
	res, err := utils.RequestResponseRapidApi(ctx, url, api.Configs.ApiHost.OpenWeatherMap, api.Configs.RapidApiKey)
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

//STUBS

// func (ws *OpenWeatherResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
// 	//save response to file
// 	file, err := ioutil.ReadFile("./stubsJSON/responseOpenWeatherMap")
// 	if err != nil {
// 		return utils.Data{}, err
// 	}

// 	var d OpenWeatherResponseStruct

// 	_ = json.Unmarshal([]byte(file), &d)

// 	stdData := utils.Data{
// 		Name:        "OpenWeatherMapApi",
// 		Location:    d.Name,
// 		Temperature: d.Main.Celsius,
// 	}

// 	return stdData, nil
// }
