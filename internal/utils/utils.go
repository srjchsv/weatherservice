package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	configs = LoadApiConfig(getEnv("APP_APIKEY_PATH", "./configs/.configs"))

	APIKEY = configs.RapidApiKey

	URLyahoo          = configs.Url.YahooWeather
	URLopenWeatherMap = configs.Url.OpenWeatherMap
	URLweatherApi     = configs.Url.WeatherApi

	APIhostYahoo          = configs.ApiHost.YahooWeather
	APIhostOpenWeatherMap = configs.ApiHost.OpenWeatherMap
	APIhostWeatherApi     = configs.ApiHost.WeatherApi
)

//getEnv gets the enviroment variable
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

type apiConfigData struct {
	RapidApiKey string `json:"RapidApiKey"`
	Url         struct {
		OpenWeatherMap string `json:"OpenWeatherMap"`
		YahooWeather   string `json:"YahooWeather"`
		WeatherApi     string `json:"WeatherApi"`
	} `json:"Url"`
	ApiHost struct {
		OpenWeatherMap string `json:"OpenWeatherMap"`
		YahooWeather   string `json:"YahooWeather"`
		WeatherApi     string `json:"WeatherApi"`
	} `json:"ApiHost"`
}

type Data struct {
	Name        string
	Location    string
	Temperature float64
}

//RequestResponseRapidApi sends http request and get response from rapid api
func RequestResponseRapidApi(ctx *gin.Context, url, apiHost string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("X-RapidAPI-Key", APIKEY)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return res, nil
}

//LoadApiConfig loads api configs
func LoadApiConfig(filename string) apiConfigData {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(2)
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		os.Exit(2)
	}
	return c
}

func printWord() {
	fmt.Println("ok")
}
