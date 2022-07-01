package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	configs = LoadApiConfig("./configs/.configs")
	APIKEY  = configs.RapidApiKey
)

type apiConfigData struct {
	RapidApiKey string `json:"RapidApiKey"`
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
