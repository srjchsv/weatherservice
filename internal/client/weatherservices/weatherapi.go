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

	url := ws.ApiUrl + "?q=" + url.QueryEscape(location)
	res, err := utils.RequestResponseRapidApi(ctx, url, ws.ApiHost, ws.ApiKey, ws.Client)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d WeatherApiResponse

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
