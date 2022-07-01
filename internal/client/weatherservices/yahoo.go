package weatherservices

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/pkg/utils"
)

type YahooResponseStruct struct {
	Location struct {
		Name string `json:"city"`
	} `json:"location"`
	Main struct {
		Condition struct {
			Celsius float64 `json:"temperature"`
		} `json:"condition"`
	} `json:"current_observation"`
}

func YahooWeatherApi(ctx *gin.Context, location string) (utils.Data, error) {

	url := "https://yahoo-weather5.p.rapidapi.com/weather?location=" + url.QueryEscape(location) + "&format=json&u=c"
	apiHost := "yahoo-weather5.p.rapidapi.com"

	res, err := utils.RequestResponseRapidApi(ctx, url, apiHost)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d YahooResponseStruct

	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return utils.Data{}, err
	}

	stdData := utils.Data{
		Name:        "YahooRapidApi",
		Location:    d.Location.Name,
		Temperature: d.Main.Condition.Celsius,
	}

	return stdData, nil

}

////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============

func YahooWeatherApiSTUB(ctx *gin.Context, location string) (utils.Data, error) {
	file, err := ioutil.ReadFile("./stubsJSON/responseYahoo")
	if err != nil {
		return utils.Data{}, err
	}

	var d YahooResponseStruct

	_ = json.Unmarshal([]byte(file), &d)

	stdData := utils.Data{
		Name:        "YahooRapidApi",
		Location:    d.Location.Name,
		Temperature: d.Main.Condition.Celsius,
	}

	return stdData, nil
}
