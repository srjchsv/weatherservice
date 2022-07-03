package weatherservices

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
)

var YahooApi YahooResponseStruct

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

// func (ws *YahooResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
// 	url := utils.URLyahoo + "?location=" + url.QueryEscape(location) + "&format=json&u=c"
// 	apiHost := utils.APIhostYahoo

// 	res, err := utils.RequestResponseRapidApi(ctx, url, apiHost)
// 	if err != nil {
// 		return utils.Data{}, err
// 	}
// 	defer res.Body.Close()

// 	var d YahooResponseStruct

// 	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
// 		return utils.Data{}, err
// 	}

// 	stdData := utils.Data{
// 		Name:        "YahooRapidApi",
// 		Location:    d.Location.Name,
// 		Temperature: d.Main.Condition.Celsius,
// 	}

// 	return stdData, nil
// }

//STUB

func (ws *YahooResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	file, err := ioutil.ReadFile("./stubsJSON/responseYahoo")
	if err != nil {
		return utils.Data{}, err
	}

	var d YahooResponseStruct

	_ = json.Unmarshal(file, &d)

	stdData := utils.Data{
		Name:        "YahooRapidApi",
		Location:    d.Location.Name,
		Temperature: d.Main.Condition.Celsius,
	}

	return stdData, nil
}
