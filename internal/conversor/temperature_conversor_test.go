package conversor

import (
	"testing"

	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertWeatherResponse_ZeroCelsius(t *testing.T) {
	// Arrange
	weather := model.WeatherResponse{}
	weather.Current.TempC = 0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, 0.0, result.Celsius)
	assert.Equal(t, 32.0, result.Fahrenheit)
	assert.Equal(t, 273.15, result.Kelvin)
}

func TestConvertWeatherResponse_PositiveCelsius(t *testing.T) {
	// Arrange
	weather := model.WeatherResponse{}
	weather.Current.TempC = 25.0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, 25.0, result.Celsius)
	assert.Equal(t, 77.0, result.Fahrenheit)
	assert.Equal(t, 298.15, result.Kelvin)
}

func TestConvertWeatherResponse_NegativeCelsius(t *testing.T) {
	// Arrange
	weather := model.WeatherResponse{}
	weather.Current.TempC = -10.0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, -10.0, result.Celsius)
	assert.Equal(t, 14.0, result.Fahrenheit)
	assert.Equal(t, 263.15, result.Kelvin)
}

func TestConvertWeatherResponse_BoilingPointWater(t *testing.T) {
	// Arrange - 100°C (boiling point of water)
	weather := model.WeatherResponse{}
	weather.Current.TempC = 100.0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, 100.0, result.Celsius)
	assert.Equal(t, 212.0, result.Fahrenheit)
	assert.Equal(t, 373.15, result.Kelvin)
}

func TestConvertWeatherResponse_AbsoluteZero(t *testing.T) {
	// Arrange - -273.15°C (absolute zero)
	weather := model.WeatherResponse{}
	weather.Current.TempC = -273.15

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, -273.15, result.Celsius)
	assert.InDelta(t, -459.67, result.Fahrenheit, 0.01) // Using InDelta for floating point comparison
	assert.InDelta(t, 0.0, result.Kelvin, 0.01)
}

func TestConvertWeatherResponse_DecimalValues(t *testing.T) {
	// Arrange
	weather := model.WeatherResponse{}
	weather.Current.TempC = 28.5

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, 28.5, result.Celsius)
	assert.InDelta(t, 83.3, result.Fahrenheit, 0.01) // Using InDelta for floating point comparison
	assert.Equal(t, 301.65, result.Kelvin)
}

func TestConvertWeatherResponse_TypicalSummerDay(t *testing.T) {
	// Arrange - Typical hot summer day
	weather := model.WeatherResponse{}
	weather.Current.TempC = 35.0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, 35.0, result.Celsius)
	assert.Equal(t, 95.0, result.Fahrenheit)
	assert.Equal(t, 308.15, result.Kelvin)
}

func TestConvertWeatherResponse_TypicalWinterDay(t *testing.T) {
	// Arrange - Typical cold winter day
	weather := model.WeatherResponse{}
	weather.Current.TempC = -5.0

	// Act
	result := ConvertWeatherResponse(weather)

	// Assert
	assert.Equal(t, -5.0, result.Celsius)
	assert.Equal(t, 23.0, result.Fahrenheit)
	assert.Equal(t, 268.15, result.Kelvin)
}
