package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/api"
)

func GetLocation(ctx *gin.Context) {
	r, err := regexp.Compile(`^[a-zA-Z',.\s-]{1,25}$`)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	location := ctx.Query("location")

	if len(location) == 0 {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	if !r.MatchString(location) {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	if r.MatchString(location) {
		weatherData, err := api.WeatherService(ctx, location)
		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"weatherData": weatherData,
			"location":    location,
		})
	}
}
