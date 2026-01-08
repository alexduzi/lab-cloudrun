package conversor

import "github.com/alexduzi/labcloudrun/internal/model"

func ConvertWeatherResponse(weather model.WeatherResponse) model.TemperatureResponse {
	kelvin := weather.Current.TempC + 273.15
	fahrenheit := weather.Current.TempC*1.8 + 32
	return model.TemperatureResponse{
		Celsius:    weather.Current.TempC,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}
}
