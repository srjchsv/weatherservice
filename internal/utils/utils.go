package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ApiConfigData struct {
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

//GetEnv gets the enviroment variable
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

//RequestResponseRapidApi sends http request and get response from rapid api
func RequestResponseRapidApi(ctx *gin.Context, url, apiHost, apiKey string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx.Request.Context(), "GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return res, nil
}

//LoadApiConfig loads api configs
func LoadApiConfig(bytes []byte) ApiConfigData {
	var c ApiConfigData
	err := json.Unmarshal(bytes, &c)
	if err != nil {
		log.Info(err)
	}
	return c
}

type Provider struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Host string `json:"host"`
}

type Providers struct {
	Key       string     `json:"key"`
	Providers []Provider `json:"services"`
}

func LoadConfig(bytes []byte) Providers {
	var providers Providers
	err := json.Unmarshal(bytes, &providers)
	if err != nil {
		log.Info(err)
	}
	return providers
}
