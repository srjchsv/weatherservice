package api

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/utils"
	"golang.org/x/sync/errgroup"
)

var (
	LocationRegexp = regexp.MustCompile(`^[a-zA-Z',.\s-]{1,25}$`)
)

type WeatherService struct {
	mu       sync.RWMutex
	Services []WeatherServiceApis
	Timeout  time.Time
}

type WeatherServiceApis interface {
	GetWeather(ctx *gin.Context, location string) (utils.Data, error)
}

func NewWeatherService(services []WeatherServiceApis) (*WeatherService, error) {
	if len(services) == 0 {
		return nil, errors.New("error empty services slice")
	}

	return &WeatherService{
		mu:       sync.RWMutex{},
		Services: services,
	}, nil
}

//Handle gets weather data from all available services
func (ws *WeatherService) Handle(ctx *gin.Context) {
	location := ctx.Query("location")

	if len(location) == 0 {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	if !LocationRegexp.MatchString(location) {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	allData := []utils.Data{}

	var eg errgroup.Group

	for _, api := range ws.Services {
		goApi := api
		eg.Go(func() error {
			return func(api WeatherServiceApis, ctx *gin.Context, location string) error {
				data, err := api.GetWeather(ctx, location)
				if err != nil {
					return err
				}
				// Lock and then defer unlock
				ws.mu.Lock()
				defer ws.mu.Unlock()
				allData = append(allData, data)
				// or mu.unlock
				return nil
			}(goApi, ctx, location)
		})
	}
	if err := eg.Wait(); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintln(err))
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"weatherData": allData,
		"location":    location,
	})
}
