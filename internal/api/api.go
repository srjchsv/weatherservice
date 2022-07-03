package api

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/client/weatherservices"
	"github.com/srjchsv/weatherservice/internal/utils"
	"golang.org/x/sync/errgroup"
)

var (
	mu sync.RWMutex
	eg errgroup.Group
)

type WeatherServiceApis interface {
	GetWeather(ctx *gin.Context, location string) (utils.Data, error)
}

//WeatherService gets wether data from all available services
func WeatherService(ctx *gin.Context, location string) ([]utils.Data, error) {
	weatherChan := []utils.Data{}

	services := []WeatherServiceApis{
		&weatherservices.OpenWeatherMapApi,
		&weatherservices.YahooApi,
		&weatherservices.WeatherApi,
	}

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
				weatherChan = append(weatherChan, data)

				return nil
			}(goApi, ctx, location)
		})
	}
	if err := eg.Wait(); err != nil {
		return weatherChan, err
	}

	return weatherChan, nil
}

func printWord() {
	fmt.Println("ok")
}
