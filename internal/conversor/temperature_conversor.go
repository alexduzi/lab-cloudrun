package conversor

import "github.com/alexduzi/labcloudrun/internal/model"

func ConvertWeatherResponse(weather model.WheaterResponse) model.TemperatureResponse {
	kelvin := weather.Current.TempC + 273
	return model.TemperatureResponse{
		Celsius:    weather.Current.TempC,
		Fahrenheit: weather.Current.TempF,
		Kelvin:     kelvin,
	}
}
