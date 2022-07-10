package router

import (
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/api"
	"github.com/thinkerou/favicon"
)

func RegisterRouter(weatherService *api.WeatherService) *gin.Engine {
	r := gin.Default()
	r.Static("/styles", "./templates/styles")
	r.Use(favicon.New("./templates/resources/favicon.ico"))
	r.LoadHTMLGlob("templates/*.html")

	home := r.Group("/")
	{
		home.GET("/", weatherService.Handle)
	}

	return r
}
