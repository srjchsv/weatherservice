package api

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/client/weatherservices"
	"github.com/srjchsv/weatherservice/pkg/utils"
)

var (
	wg sync.WaitGroup
	mu sync.RWMutex
)

func WeatherService(ctx *gin.Context, location string) ([]utils.Data, error) {
	weatherChan := []utils.Data{}

	services := []func(ctx *gin.Context, location string) (utils.Data, error){
		weatherservices.WeatherApi,
		weatherservices.YahooWeatherApi,
		weatherservices.OpenWeatherMapApi,
	}

	for _, api := range services {
		wg.Add(1)
		go func(api func(ctx *gin.Context, location string) (utils.Data, error)) error {
			defer wg.Done()
			defer mu.Unlock()
			data, err := api(ctx, location)
			if err != nil {
				return err
			}
			mu.Lock()
			weatherChan = append(weatherChan, data)

			return nil
		}(api)
	}
	wg.Wait()
	return weatherChan, nil
}
