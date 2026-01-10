package model

import "time"

// ViacepResponse represents the response from ViaCEP API
type ViacepResponse struct {
	Erro        *string `json:"erro,omitempty" example:"true"`
	Cep         string  `json:"cep" example:"01310-100"`
	Logradouro  string  `json:"logradouro" example:"Avenida Paulista"`
	Complemento string  `json:"complemento" example:"de 612 a 1510 - lado par"`
	Unidade     string  `json:"unidade" example:""`
	Bairro      string  `json:"bairro" example:"Bela Vista"`
	Localidade  string  `json:"localidade" example:"São Paulo"`
	Uf          string  `json:"uf" example:"SP"`
	Estado      string  `json:"estado" example:"São Paulo"`
	Regiao      string  `json:"regiao" example:"Sudeste"`
	Ibge        string  `json:"ibge" example:"3550308"`
	Gia         string  `json:"gia" example:"1004"`
	Ddd         string  `json:"ddd" example:"11"`
	Siafi       string  `json:"siafi" example:"7107"`
}

type WeatherResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
		ShortRad   float64 `json:"short_rad"`
		DiffRad    float64 `json:"diff_rad"`
		Dni        float64 `json:"dni"`
		Gti        float64 `json:"gti"`
	} `json:"current"`
}

// TemperatureResponse represents temperature in different units
type TemperatureResponse struct {
	Celsius    float64 `json:"temp_C" example:"28.5"`
	Fahrenheit float64 `json:"temp_F" example:"83.3"`
	Kelvin     float64 `json:"temp_K" example:"301.65"`
}

// StatusResponse represents the health/readiness status response
type StatusResponse struct {
	Status    string    `json:"status" example:"healthy"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	Service   string    `json:"service" example:"lab-cloudrun-api"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message" example:"invalid zipcode"`
}
