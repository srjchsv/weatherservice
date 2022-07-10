package weatherservices

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
)

type YahooApi struct {
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

func NewYahooApi(apiHost, apiKey, apiUrl, ApiName string, httpClient *http.Client, response interface{}) *YahooApi {
	return &YahooApi{
		ApiHost:  apiHost,
		ApiKey:   apiKey,
		ApiUrl:   apiUrl,
		ApiName:  ApiName,
		Client:   httpClient,
		Response: response,
	}
}

func (ws *YahooApi) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	url := ws.ApiUrl + "?location=" + url.QueryEscape(location) + "&format=json&u=c"
	res, err := utils.RequestResponseRapidApi(ctx, url, ws.ApiHost, ws.ApiKey, ws.Client)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d YahooResponse

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

// //STUB

// // func (ws *YahooResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
// // 	file, err := ioutil.ReadFile("./stubsJSON/responseYahoo")
// // 	if err != nil {
// // 		return utils.Data{}, err
// // 	}

// // 	var d YahooResponseStruct

// // 	_ = json.Unmarshal(file, &d)

// // 	stdData := utils.Data{
// // 		Name:        "YahooRapidApi",
// // 		Location:    d.Location.Name,
// // 		Temperature: d.Main.Condition.Celsius,
// // 	}

// // 	return stdData, nil
// // }
