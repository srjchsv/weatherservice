package weatherservices

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
)

type WeatherApi struct {
	ApiHost  string
	ApiKey   string
	ApiUrl   string
	ApiName  string
	Client   *http.Client
	Response interface{}
}

type YahooResponse struct {
	Location struct {
		Name string `json:"city"`
	} `json:"location"`
	Main struct {
		Condition struct {
			Celsius float64 `json:"temperature"`
		} `json:"condition"`
	} `json:"current_observation"`
}

type OpenWeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

type WeatherApiResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Celsius float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeatherApi(apiHost, apiKey, apiUrl, ApiName string, httpClient *http.Client, response interface{}) *WeatherApi {
	return &WeatherApi{
		ApiHost:  apiHost,
		ApiKey:   apiKey,
		ApiUrl:   apiUrl,
		ApiName:  ApiName,
		Client:   httpClient,
		Response: response,
	}
}

func (ws *WeatherApi) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	var stdData utils.Data

	switch t := ws.Response.(type) {
	case WeatherApiResponse:
		ws.ApiUrl = ws.ApiUrl + "?q=" + url.QueryEscape(location)
		res, err := utils.RequestResponseRapidApi(ctx, ws.ApiUrl, ws.ApiHost, ws.ApiKey, ws.Client)
		if err != nil {
			return utils.Data{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			return utils.Data{}, err
		}
		stdData = utils.Data{
			Name:        "WeatherApi",
			Location:    t.Location.Name,
			Temperature: t.Current.Celsius,
		}
	case OpenWeatherResponse:
		ws.ApiUrl = ws.ApiUrl + "?q=" + url.QueryEscape(location) + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"
		res, err := utils.RequestResponseRapidApi(ctx, ws.ApiUrl, ws.ApiHost, ws.ApiKey, ws.Client)
		if err != nil {
			return utils.Data{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			return utils.Data{}, err
		}
		stdData = utils.Data{
			Name:        "OpenWeatherMapApi",
			Location:    t.Name,
			Temperature: t.Main.Celsius,
		}
	case YahooResponse:
		ws.ApiUrl = ws.ApiUrl + "?location=" + url.QueryEscape(location) + "&format=json&u=c"
		res, err := utils.RequestResponseRapidApi(ctx, ws.ApiUrl, ws.ApiHost, ws.ApiKey, ws.Client)
		if err != nil {
			return utils.Data{}, err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			return utils.Data{}, err
		}
		stdData = utils.Data{
			Name:        "YahooRapidApi",
			Location:    t.Location.Name,
			Temperature: t.Main.Condition.Celsius,
		}
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
