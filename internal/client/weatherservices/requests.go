package weatherservices

import (
	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
)

type WeatherApiResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Celsius float64 `json:"temp_c"`
	} `json:"current"`
}

func (ws *WeatherApi) weatherApiRequest(ctx *gin.Context, location string) (utils.Data, error) {
	ws.ApiUrl = ws.ApiUrl + "?q=" + url.QueryEscape(location)
	res, err := utils.RequestResponseRapidApi(ctx, ws.ApiUrl, ws.ApiHost, ws.ApiKey, ws.Client)
	if err != nil {
		return utils.Data{}, err
	}
	var t WeatherApiResponse
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return utils.Data{}, err
	}
	return utils.Data{
		Name:        "WeatherApi",
		Location:    t.Location.Name,
		Temperature: t.Current.Celsius,
	}, nil
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

func (ws *WeatherApi) yahoooApiRequest(ctx *gin.Context, location string) (utils.Data, error) {
	ws.ApiUrl = ws.ApiUrl + "?location=" + url.QueryEscape(location) + "&format=json&u=c"
	res, err := utils.RequestResponseRapidApi(ctx, ws.ApiUrl, ws.ApiHost, ws.ApiKey, ws.Client)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()
	var t YahooResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return utils.Data{}, err
	}
	return utils.Data{
		Name:        "YahooRapidApi",
		Location:    t.Location.Name,
		Temperature: t.Main.Condition.Celsius,
	}, nil
}
