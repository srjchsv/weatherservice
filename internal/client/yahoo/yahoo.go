package yahoo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	apiConfig, err := utils.LoadApiConfig(ctx, "./configs/.configs")
	if err != nil {
		return utils.Data{}, err
	}

	url := "https://yahoo-weather5.p.rapidapi.com/weather?location=" + location + "&format=json&u=c"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.Data{}, err
	}

	req.Header.Add("X-RapidAPI-Key", apiConfig.OpenWeatherMapApiKey)
	req.Header.Add("X-RapidAPI-Host", "yahoo-weather5.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
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
