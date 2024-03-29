package utils

import (
	"encoding/json"
	"net/http"

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

// RequestResponseRapidApi sends http request and get response from rapid api
func RequestResponseRapidApi(ctx *gin.Context, url, apiHost, apiKey string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LoadConfig parses JSON config into a struct
func LoadConfig(bytes []byte) Providers {
	var providers Providers
	err := json.Unmarshal(bytes, &providers)
	if err != nil {
		log.Info(err)
	}
	return providers
}
