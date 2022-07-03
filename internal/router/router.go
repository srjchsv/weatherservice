package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/handlers"
	"github.com/thinkerou/favicon"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/styles", "./templates/styles")
	r.Use(favicon.New("./templates/resources/favicon.ico"))
	r.LoadHTMLGlob("templates/*.html")

	home := r.Group("/")
	{
		home.GET("/", handlers.GetLocation)
	}

	return r
}

func printWord() {
	fmt.Println("ok")
}
