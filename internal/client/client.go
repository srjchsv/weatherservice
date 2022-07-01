package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

// func QueryOpenWeatherMapApi(ctx *gin.Context, location string) (weatherData, error) {
// 	apiConfig := utils.LoadApiConfig(ctx, "./configs/.configs")

// 	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=" + location + "&lat=0&lon=0&id=2172797&lang=null&units=metric&mode=json"

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return weatherData{}, err
// 	}
// 	req.Header.Add("X-RapidAPI-Key", apiConfig.OpenWeatherMapApiKey)
// 	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return weatherData{}, err
// 	}
// 	defer res.Body.Close()

// 	// // //save response to file

// 	// file, err := os.OpenFile("response", os.O_CREATE|os.O_WRONLY, 0666)
// 	// if err != nil {
// 	// 	return weatherData{}, err
// 	// }
// 	// defer file.Close()

// 	// io.Copy(file, res.Body)

// 	// // //// end save response to txt

// 	var d weatherData

// 	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
// 		return weatherData{}, err
// 	}

// 	return d, nil
// }

////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============
////++=============////++=============

func QueryOpenWeatherMapApi(ctx *gin.Context, location string) (weatherData, error) {

	//save response to file
	file, err := ioutil.ReadFile("response.json")
	if err != nil {
		return weatherData{}, err
	}

	var d weatherData

	_ = json.Unmarshal([]byte(file), &d)

	return d, nil
}
