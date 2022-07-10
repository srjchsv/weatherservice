package utils

import (
	"os"
	"reflect"
	"testing"
)

var bytes = []byte(`{
    "RapidApiKey" : "apikey",
    "Url" : {
        "OpenWeatherMap": "https://community-open-weather-map.p.rapidapi.com/weather" ,
        "YahooWeather":"https://yahoo-weather5.p.rapidapi.com/weather" ,
        "WeatherApi" :"https://weatherapi-com.p.rapidapi.com/current.json"
    },
    "ApiHost" : {
        "OpenWeatherMap":  "community-open-weather-map.p.rapidapi.com",
        "YahooWeather":"yahoo-weather5.p.rapidapi.com",
        "WeatherApi" : "weatherapi-com.p.rapidapi.com"
    }
}
`)

func TestGetEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "one",
			args: args{
				key:      "",
				fallback: "fallback",
			},
			want: "fallback",
		},
		{
			name: "two",
			args: args{
				key:      "key",
				fallback: "fallback",
			},
			want: "key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.args.key) != 0 {
				os.Setenv("key", "key")
			}
			if got := GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadApiConfig(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want ApiConfigData
	}{
		// TODO: Add test cases.
		{name: "one", args: args{bytes: bytes}, want: ApiConfigData{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadApiConfig(tt.args.bytes)

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("LoadApiConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
