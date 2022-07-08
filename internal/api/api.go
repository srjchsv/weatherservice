package api

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
	"golang.org/x/sync/errgroup"
)

var (
	mu       sync.RWMutex
	eg       errgroup.Group
	Configs  utils.ApiConfigData
	services []WeatherServiceApis
	
	Services = func(cfg utils.ApiConfigData, arr []WeatherServiceApis) {
		services = arr
		Configs = cfg
	}
)

type WeatherServiceApis interface {
	GetWeather(ctx *gin.Context, location string) (utils.Data, error)
}

//WeatherService gets weather data from all available services
func WeatherService(ctx *gin.Context, location string) ([]utils.Data, error) {
	allData := []utils.Data{}

	for _, api := range services {
		goApi := api
		eg.Go(func() error {
			return func(api WeatherServiceApis, ctx *gin.Context, location string) error {
				defer mu.Unlock()
				data, err := api.GetWeather(ctx, location)
				if err != nil {
					return err
				}
				mu.Lock()
				allData = append(allData, data)

				return nil
			}(goApi, ctx, location)
		})
	}
	if err := eg.Wait(); err != nil {
		return allData, err
	}
	return allData, nil
}
