package conversor

import (
	"math"

	"github.com/alexduzi/labcloudrun/internal/model"
)

func ConvertWeatherResponse(weather model.WeatherResponse) model.TemperatureResponse {
	kelvin := weather.Current.TempC + 273.15
	fahrenheit := weather.Current.TempC*1.8 + 32
	return model.TemperatureResponse{
		Celsius:    roundToTwoDecimals(weather.Current.TempC),
		Fahrenheit: roundToTwoDecimals(fahrenheit),
		Kelvin:     roundToTwoDecimals(kelvin),
	}
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
