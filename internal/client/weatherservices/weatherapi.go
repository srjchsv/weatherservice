package weatherservices

import (
	"net/http"

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

	switch ws.Response.(type) {
	case WeatherApiResponse:
		stData, err := ws.weatherApiRequest(ctx, location)
		return stData, err
	case YahooResponse:
		stData, err := ws.yahoooApiRequest(ctx, location)
		return stData, err
	}
	return stdData, nil
}
