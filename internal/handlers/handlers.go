package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/weatherservice/internal/weatherservice"
)

func GetLocation(ctx *gin.Context) {
	r, err := regexp.Compile(`^[a-zA-z]{0,20}$`)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	location := ctx.Query("location")

	if r.MatchString(location) {
		weatherData, err := weatherservice.WeatherService(ctx, location)
		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintln(err))
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"weatherData":   weatherData,
			"locationCheck": location,
		})

	} else {
		ctx.String(http.StatusBadRequest, "Bad request")
	}
}
