package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/srjchsv/weatherservice/internal/api"
	"github.com/srjchsv/weatherservice/internal/client/weatherservices"
	"github.com/srjchsv/weatherservice/internal/router"
	"github.com/srjchsv/weatherservice/internal/utils"
)

func main() {
	//Read configs
	bytes, err := ioutil.ReadFile(utils.GetEnv("APP_APIKEY_PATH", "./configs/.services"))
	if err != nil {
		log.Info(err)
	}
	srv := utils.LoadConfig(bytes)
	//cfg := utils.LoadApiConfig(bytes)

	// m := make(map[string]interface{})
	// m["OpenWeatherMap"] = weatherservices.OpenWeatherResponse{}
	// m["YahooWeather"] = weatherservices.YahooResponse{}
	// m["WeatherApi"] = weatherservices.WeatherApiResponse{}

	// var apis []api.WeatherServiceApis

	// for _, val := range srv.Providers {
	// 	srvs := weatherservices.NewWeatherApi(val.Host, srv.Key, val.Url, val.Name, client, m[val.Name])
	// 	apis = append(apis, srvs)
	// }

	client := &http.Client{Timeout: time.Second * 5}

	wa := weatherservices.NewWeatherApi(srv.Providers[2].Host, srv.Key, srv.Providers[2].Url, srv.Providers[2].Name, client, weatherservices.WeatherApiResponse{})
	yw := weatherservices.NewYahooApi(srv.Providers[1].Host, srv.Key, srv.Providers[1].Url, srv.Providers[1].Name, client, weatherservices.YahooResponse{})
	ow := weatherservices.NewOpenWeatherApi(srv.Providers[0].Host, srv.Key, srv.Providers[0].Url, srv.Providers[0].Name, client, weatherservices.OpenWeatherResponse{})

	apis := []api.WeatherServiceApis{wa, yw, ow}

	weatherService, err := api.NewWeatherService(apis)
	if err != nil {
		log.Fatal(err)
	}
	r := router.RegisterRouter(weatherService)
	r.Run(":8080")
}
