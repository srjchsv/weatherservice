package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/srjchsv/weatherservice/internal/client/weatherservices"
	"github.com/srjchsv/weatherservice/internal/handlers"
	"github.com/srjchsv/weatherservice/internal/router"
	"github.com/srjchsv/weatherservice/internal/utils"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		log.Fatal("error no api key.")
	}
	//Read configs
	bytes, err := ioutil.ReadFile("./configs/.configs")
	if err != nil {
		log.Info(err)
	}
	//Load configs
	cfg := utils.LoadConfig(bytes)
	//Client configs
	client := &http.Client{Timeout: time.Second * 5}

	//Map of response structs for services
	m := make(map[string]interface{})
	m["YahooWeather"] = weatherservices.YahooResponse{}
	m["WeatherApi"] = weatherservices.WeatherApiResponse{}

	var apis []handlers.WeatherServiceApis

	//Initialize services
	for _, val := range cfg.Providers {
		api := weatherservices.NewWeatherApi(val.Host, apiKey, val.Url, val.Name, client, m[val.Name])
		apis = append(apis, api)
	}
	//Pass services to weather service handler
	weatherService, err := handlers.NewWeatherService(apis)
	if err != nil {
		log.Fatal(err)
	}
	//register router and run server
	r := router.RegisterRouter(weatherService)
	r.Run(":8080")
}
