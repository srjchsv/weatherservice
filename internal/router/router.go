package router

import (
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/handlers"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/styles", "./templates/styles")
	r.StaticFile("favicon.ico", "./templates/resources/favicon.ico")

	home := r.Group("/")
	{
		home.GET("/", handlers.GetLocation)
	}

	return r
}
