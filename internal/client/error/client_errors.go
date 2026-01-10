package error

import (
	"errors"
	"fmt"
)

var (
	CepClientBadRequest      = errors.New("invalid request to CEP API")
	CepClientNotFound        = errors.New("CEP API returned not found")
	CepClientInternalError   = errors.New("CEP API internal error")
	CepClientUnexpectedError = errors.New("unexpected error from CEP API")

	WeatherClientBadRequest      = errors.New("invalid request to Weather API")
	WeatherClientNotFound        = errors.New("Weather API returned not found")
	WeatherClientInternalError   = errors.New("Weather API internal error")
	WeatherClientUnexpectedError = errors.New("unexpected error from Weather API")
)

func NewCepClientHTTPError(statusCode int) error {
	switch statusCode {
	case 400:
		return CepClientBadRequest
	case 404:
		return CepClientNotFound
	case 500, 502, 503, 504:
		return CepClientInternalError
	default:
		return fmt.Errorf("%w: status code %d", CepClientUnexpectedError, statusCode)
	}
}

func NewWeatherClientHTTPError(statusCode int) error {
	switch statusCode {
	case 400:
		return WeatherClientBadRequest
	case 404:
		return WeatherClientNotFound
	case 500, 502, 503, 504:
		return WeatherClientInternalError
	default:
		return fmt.Errorf("%w: status code %d", WeatherClientUnexpectedError, statusCode)
	}
}
