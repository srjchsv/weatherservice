package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Data struct {
	Name        string
	Location    string
	Temperature float64
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

//LoadConfig parses JSON config into a struct
func LoadConfig(bytes []byte) Providers {
	var providers Providers
	err := json.Unmarshal(bytes, &providers)
	if err != nil {
		log.Info(err)
	}
	return providers
}
