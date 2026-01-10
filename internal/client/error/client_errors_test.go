package error

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCepClientHTTPError_BadRequest(t *testing.T) {
	// act
	err := NewCepClientHTTPError(400)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientBadRequest)
	assert.Equal(t, "invalid request to CEP API", err.Error())
}

func TestNewCepClientHTTPError_NotFound(t *testing.T) {
	// act
	err := NewCepClientHTTPError(404)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientNotFound)
	assert.Equal(t, "CEP API returned not found", err.Error())
}

func TestNewCepClientHTTPError_InternalServerError(t *testing.T) {
	// act
	err := NewCepClientHTTPError(500)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientInternalError)
	assert.Equal(t, "CEP API internal error", err.Error())
}

func TestNewCepClientHTTPError_BadGateway(t *testing.T) {
	// act
	err := NewCepClientHTTPError(502)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientInternalError)
	assert.Equal(t, "CEP API internal error", err.Error())
}

func TestNewCepClientHTTPError_ServiceUnavailable(t *testing.T) {
	// act
	err := NewCepClientHTTPError(503)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientInternalError)
	assert.Equal(t, "CEP API internal error", err.Error())
}

func TestNewCepClientHTTPError_GatewayTimeout(t *testing.T) {
	// act
	err := NewCepClientHTTPError(504)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientInternalError)
	assert.Equal(t, "CEP API internal error", err.Error())
}

func TestNewCepClientHTTPError_UnexpectedStatusCode(t *testing.T) {
	// act
	err := NewCepClientHTTPError(418)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientUnexpectedError)
	assert.Contains(t, err.Error(), "unexpected error from CEP API")
	assert.Contains(t, err.Error(), "status code 418")
}

func TestNewCepClientHTTPError_AnotherUnexpectedStatusCode(t *testing.T) {
	// act
	err := NewCepClientHTTPError(429)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, CepClientUnexpectedError)
	assert.Contains(t, err.Error(), "unexpected error from CEP API")
	assert.Contains(t, err.Error(), "status code 429")
}

func TestNewWeatherClientHTTPError_BadRequest(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(400)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientBadRequest)
	assert.Equal(t, "invalid request to Weather API", err.Error())
}

func TestNewWeatherClientHTTPError_NotFound(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(404)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientNotFound)
	assert.Equal(t, "Weather API returned not found", err.Error())
}

func TestNewWeatherClientHTTPError_InternalServerError(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(500)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientInternalError)
	assert.Equal(t, "Weather API internal error", err.Error())
}

func TestNewWeatherClientHTTPError_BadGateway(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(502)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientInternalError)
	assert.Equal(t, "Weather API internal error", err.Error())
}

func TestNewWeatherClientHTTPError_ServiceUnavailable(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(503)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientInternalError)
	assert.Equal(t, "Weather API internal error", err.Error())
}

func TestNewWeatherClientHTTPError_GatewayTimeout(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(504)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientInternalError)
	assert.Equal(t, "Weather API internal error", err.Error())
}

func TestNewWeatherClientHTTPError_UnexpectedStatusCode(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(418)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientUnexpectedError)
	assert.Contains(t, err.Error(), "unexpected error from Weather API")
	assert.Contains(t, err.Error(), "status code 418")
}

func TestNewWeatherClientHTTPError_AnotherUnexpectedStatusCode(t *testing.T) {
	// act
	err := NewWeatherClientHTTPError(429)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, WeatherClientUnexpectedError)
	assert.Contains(t, err.Error(), "unexpected error from Weather API")
	assert.Contains(t, err.Error(), "status code 429")
}

func TestPredefinedErrors_AreDistinct(t *testing.T) {
	// assert - verificar que todos os erros são distintos
	assert.NotEqual(t, CepClientBadRequest, CepClientNotFound)
	assert.NotEqual(t, CepClientBadRequest, CepClientInternalError)
	assert.NotEqual(t, CepClientBadRequest, CepClientUnexpectedError)
	assert.NotEqual(t, CepClientNotFound, CepClientInternalError)
	assert.NotEqual(t, CepClientNotFound, CepClientUnexpectedError)
	assert.NotEqual(t, CepClientInternalError, CepClientUnexpectedError)

	assert.NotEqual(t, WeatherClientBadRequest, WeatherClientNotFound)
	assert.NotEqual(t, WeatherClientBadRequest, WeatherClientInternalError)
	assert.NotEqual(t, WeatherClientBadRequest, WeatherClientUnexpectedError)
	assert.NotEqual(t, WeatherClientNotFound, WeatherClientInternalError)
	assert.NotEqual(t, WeatherClientNotFound, WeatherClientUnexpectedError)
	assert.NotEqual(t, WeatherClientInternalError, WeatherClientUnexpectedError)
}

func TestErrorIs_WithWrappedErrors(t *testing.T) {
	// arrange
	cepErr := NewCepClientHTTPError(418)
	weatherErr := NewWeatherClientHTTPError(429)

	// assert - verificar que errors.Is funciona corretamente com erros wrapped
	assert.True(t, errors.Is(cepErr, CepClientUnexpectedError))
	assert.False(t, errors.Is(cepErr, CepClientBadRequest))
	assert.False(t, errors.Is(cepErr, WeatherClientUnexpectedError))

	assert.True(t, errors.Is(weatherErr, WeatherClientUnexpectedError))
	assert.False(t, errors.Is(weatherErr, WeatherClientBadRequest))
	assert.False(t, errors.Is(weatherErr, CepClientUnexpectedError))
}
