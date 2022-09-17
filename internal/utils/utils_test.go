package utils

import (
	"reflect"
	"testing"
)

var bytes = []byte(`{
    "key": "77777777777777777777",
    "services": [
        {
            "name": "OpenWeatherMap",
            "url": "https://community-open-weather-map.p.rapidapi.com/weather",
            "host": "community-open-weather-map.p.rapidapi.com"
        },
        {
            "name": "YahooWeather",
            "url": "https://yahoo-weather5.p.rapidapi.com/weather",
            "host": "yahoo-weather5.p.rapidapi.com"
        },
        {
            "name": "WeatherApi",
            "url": "https://weatherapi-com.p.rapidapi.com/current.json",
            "host": "weatherapi-com.p.rapidapi.com"
        }
    ]
}
`)

func TestLoadApiConfig(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want Providers
	}{
		// TODO: Add test cases.
		{name: "one", args: args{bytes: bytes}, want: Providers{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadConfig(tt.args.bytes)

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("LoadApiConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
