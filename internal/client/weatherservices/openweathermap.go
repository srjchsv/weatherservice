package weatherservices

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
)

type OpenWeatherApi struct {
	ApiHost  string
	ApiKey   string
	ApiUrl   string
	ApiName  string
	Client   *http.Client
	Response interface{}
}

type OpenWeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func NewOpenWeatherApi(apiHost, apiKey, apiUrl, ApiName string, httpClient *http.Client, response interface{}) *OpenWeatherApi {
	return &OpenWeatherApi{
		ApiHost:  apiHost,
		ApiKey:   apiKey,
		ApiUrl:   apiUrl,
		ApiName:  ApiName,
		Client:   httpClient,
		Response: response,
	}
}

func (ws *OpenWeatherApi) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
	url := ws.ApiUrl + "?q=" + url.QueryEscape(location) + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"
	res, err := utils.RequestResponseRapidApi(ctx, url, ws.ApiHost, ws.ApiKey, ws.Client)
	if err != nil {
		return utils.Data{}, err
	}
	defer res.Body.Close()

	var d OpenWeatherResponse

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

// //STUBS

// // func (ws *OpenWeatherResponseStruct) GetWeather(ctx *gin.Context, location string) (utils.Data, error) {
// // 	//save response to file
// // 	file, err := ioutil.ReadFile("./stubsJSON/responseOpenWeatherMap")
// // 	if err != nil {
// // 		return utils.Data{}, err
// // 	}

// // 	var d OpenWeatherResponseStruct

// // 	_ = json.Unmarshal([]byte(file), &d)

// // 	stdData := utils.Data{
// // 		Name:        "OpenWeatherMapApi",
// // 		Location:    d.Name,
// // 		Temperature: d.Main.Celsius,
// // 	}

// // 	return stdData, nil
// // }
