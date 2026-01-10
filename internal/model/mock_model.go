package model

func GetViacepResponseMock(zipCode string) *ViacepResponse {
	return &ViacepResponse{
		Erro:        nil,
		Cep:         zipCode,
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Unidade:     "",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Estado:      "São Paulo",
		Regiao:      "Sudeste",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}
}

func GetWeatherResponseMock(city string) *WeatherResponse {
	response := &WeatherResponse{}
	response.Location.Name = "Sao Paulo"
	response.Location.Region = "Sao Paulo"
	response.Location.Country = "Brazil"
	response.Location.Lat = -23.5333
	response.Location.Lon = -46.6167
	response.Location.TzID = "America/Sao_Paulo"
	response.Location.LocaltimeEpoch = 1768066477
	response.Location.Localtime = "2026-01-10 14:34"

	response.Current.LastUpdatedEpoch = 1768066200
	response.Current.LastUpdated = "2026-01-10 14:30"
	response.Current.TempC = 32.2
	response.Current.TempF = 90
	response.Current.IsDay = 1
	response.Current.Condition.Text = "Partly cloudy"
	response.Current.Condition.Icon = "//cdn.weatherapi.com/weather/64x64/day/116.png"
	response.Current.Condition.Code = 1003
	response.Current.WindMph = 5.4
	response.Current.WindKph = 8.6
	response.Current.WindDegree = 309
	response.Current.WindDir = "NW"
	response.Current.PressureMb = 1015
	response.Current.PressureIn = 29.97
	response.Current.PrecipMm = 0.02
	response.Current.PrecipIn = 0
	response.Current.Humidity = 36
	response.Current.Cloud = 75
	response.Current.FeelslikeC = 33.2
	response.Current.FeelslikeF = 91.8
	response.Current.WindchillC = 30.6
	response.Current.WindchillF = 87
	response.Current.HeatindexC = 30.9
	response.Current.HeatindexF = 87.6
	response.Current.DewpointC = 15.5
	response.Current.DewpointF = 59.9
	response.Current.VisKm = 10
	response.Current.VisMiles = 6
	response.Current.Uv = 11
	response.Current.GustMph = 6.6
	response.Current.GustKph = 10.7
	response.Current.ShortRad = 886
	response.Current.DiffRad = 215
	response.Current.Dni = 1602
	response.Current.Gti = 972
	return response
}
