package weatherservice

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/client/openweathermap"
	"github.com/srjchsv/weatherservice/internal/client/weatherapi"
	"github.com/srjchsv/weatherservice/internal/client/yahoo"
	"github.com/srjchsv/weatherservice/pkg/utils"
)

var (
	wg sync.WaitGroup
	mu sync.RWMutex
)

func WeatherService(ctx *gin.Context, location string) ([]utils.Data, error) {
	weatherChan := []utils.Data{}

	services := []func(ctx *gin.Context, location string) (utils.Data, error){
		weatherapi.WeatherApi,
		yahoo.YahooWeatherApi,
		openweathermap.OpenWeatherMapApi,
	}

	for _, api := range services {
		wg.Add(1)
		go func(api func(ctx *gin.Context, location string) (utils.Data, error)) error {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			data, err := api(ctx, location)
			if err != nil {
				return err
			}
			weatherChan = append(weatherChan, data)

			return nil
		}(api)
	}
	wg.Wait()
	return weatherChan, nil
}
