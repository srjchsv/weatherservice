package main

import (
	"github.com/srjchsv/weatherservice/internal/api"
	"github.com/srjchsv/weatherservice/internal/client/weatherservices"
	"github.com/srjchsv/weatherservice/internal/router"
	"github.com/srjchsv/weatherservice/internal/utils"
)

func main() {
	cfg := utils.LoadApiConfig(utils.GetEnv("APP_APIKEY_PATH", "./configs/.configs"))
	
	services := []api.WeatherServiceApis{
		&weatherservices.WeatherApi,
		&weatherservices.OpenWeatherMapApi,
		&weatherservices.YahooApi,
	}
	
	api.Services(cfg, services)
	r := router.RegisterRouter()
	r.Run(":8080")
}
