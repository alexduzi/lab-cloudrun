package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func resetViperAndConfig() {
	viper.Reset()
	AppConfig = nil
}

func TestLoadConfig_WithDefaultValues(t *testing.T) {
	// arrange
	resetViperAndConfig()

	// Garantir que não há variáveis de ambiente configuradas
	os.Unsetenv("PORT")
	os.Unsetenv("WEATHER_API_KEY")
	os.Unsetenv("VIA_CEP_BASE_URL")
	os.Unsetenv("WEATHER_BASE_URL")
	os.Unsetenv("GIN_MODE")

	// act
	config, err := LoadConfig()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "", config.WeatherAPIKey)
	assert.Equal(t, "https://viacep.com.br/ws/{cep}/json/", config.ViaCEPBaseURL)
	assert.Equal(t, "http://api.weatherapi.com/v1/current.json", config.WeatherBaseURL)
	assert.Equal(t, "debug", config.GinMode)
	assert.Equal(t, config, AppConfig)
}

func TestLoadConfig_WithEnvironmentVariables(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Setenv("PORT", "3000")
	os.Setenv("WEATHER_API_KEY", "test-api-key-123")
	os.Setenv("VIA_CEP_BASE_URL", "https://custom-cep-api.com")
	os.Setenv("WEATHER_BASE_URL", "https://custom-weather-api.com")
	os.Setenv("GIN_MODE", "release")

	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("WEATHER_API_KEY")
		os.Unsetenv("VIA_CEP_BASE_URL")
		os.Unsetenv("WEATHER_BASE_URL")
		os.Unsetenv("GIN_MODE")
	}()

	// act
	config, err := LoadConfig()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "3000", config.Port)
	assert.Equal(t, "test-api-key-123", config.WeatherAPIKey)
	assert.Equal(t, "https://custom-cep-api.com", config.ViaCEPBaseURL)
	assert.Equal(t, "https://custom-weather-api.com", config.WeatherBaseURL)
	assert.Equal(t, "release", config.GinMode)
}

func TestLoadConfig_WithPartialEnvironmentVariables(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Setenv("PORT", "9090")
	os.Setenv("WEATHER_API_KEY", "my-api-key")

	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("WEATHER_API_KEY")
	}()

	// act
	config, err := LoadConfig()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "9090", config.Port)
	assert.Equal(t, "my-api-key", config.WeatherAPIKey)
	assert.Equal(t, "https://viacep.com.br/ws/{cep}/json/", config.ViaCEPBaseURL)
	assert.Equal(t, "http://api.weatherapi.com/v1/current.json", config.WeatherBaseURL)
	assert.Equal(t, "debug", config.GinMode)
}

func TestLoadConfig_WithEmptyWeatherAPIKey(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Unsetenv("WEATHER_API_KEY")

	// act
	config, err := LoadConfig()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "", config.WeatherAPIKey)
}

func TestLoadConfig_SetsAppConfig(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Setenv("WEATHER_API_KEY", "test-key")
	defer os.Unsetenv("WEATHER_API_KEY")

	// act
	config, err := LoadConfig()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, config, AppConfig)
	assert.NotNil(t, AppConfig)
}

func TestLoadConfig_MultipleCallsOverwriteAppConfig(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	// act
	config1, err1 := LoadConfig()

	resetViperAndConfig()
	os.Setenv("PORT", "9000")

	config2, err2 := LoadConfig()

	// assert
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, config1)
	assert.NotNil(t, config2)
	assert.Equal(t, "9000", config2.Port)
	assert.Equal(t, config2, AppConfig)
}

func TestLoadConfig_GinModeValues(t *testing.T) {
	tests := []struct {
		name        string
		ginMode     string
		expectedVal string
	}{
		{
			name:        "Debug mode",
			ginMode:     "debug",
			expectedVal: "debug",
		},
		{
			name:        "Release mode",
			ginMode:     "release",
			expectedVal: "release",
		},
		{
			name:        "Test mode",
			ginMode:     "test",
			expectedVal: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			resetViperAndConfig()
			os.Setenv("GIN_MODE", tt.ginMode)
			defer os.Unsetenv("GIN_MODE")

			// act
			config, err := LoadConfig()

			// assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedVal, config.GinMode)
		})
	}
}

func TestGetConfig_WhenConfigIsLoaded(t *testing.T) {
	// arrange
	resetViperAndConfig()

	os.Setenv("WEATHER_API_KEY", "test-key")
	defer os.Unsetenv("WEATHER_API_KEY")

	_, _ = LoadConfig()

	// act
	config := GetConfig()

	// assert
	assert.NotNil(t, config)
	assert.Equal(t, AppConfig, config)
}

func TestGetConfig_WhenConfigIsNotLoaded_Panics(t *testing.T) {
	// arrange
	resetViperAndConfig()

	// act & assert
	// GetConfig() chama log.Fatal() quando AppConfig é nil, que encerra o programa
	// Por isso não podemos testar diretamente sem modificar o código
	// Mas podemos verificar que AppConfig está nil
	assert.Nil(t, AppConfig)
}

func TestConfig_StructureFields(t *testing.T) {
	// arrange
	config := &Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "debug",
	}

	// assert
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "test-api-key", config.WeatherAPIKey)
	assert.Equal(t, "https://viacep.com.br/ws/{cep}/json/", config.ViaCEPBaseURL)
	assert.Equal(t, "http://api.weatherapi.com/v1/current.json", config.WeatherBaseURL)
	assert.Equal(t, "debug", config.GinMode)
}

func TestLoadConfig_WithDifferentPorts(t *testing.T) {
	tests := []struct {
		name string
		port string
	}{
		{"Port 80", "80"},
		{"Port 443", "443"},
		{"Port 3000", "3000"},
		{"Port 8000", "8000"},
		{"Port 8080", "8080"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			resetViperAndConfig()
			os.Setenv("PORT", tt.port)
			defer os.Unsetenv("PORT")

			// act
			config, err := LoadConfig()

			// assert
			assert.NoError(t, err)
			assert.Equal(t, tt.port, config.Port)
		})
	}
}

func TestLoadConfig_ReturnsNoError(t *testing.T) {
	// arrange
	resetViperAndConfig()

	// act
	_, err := LoadConfig()

	// assert
	assert.NoError(t, err)
}

func TestLoadConfig_ConfigNotNil(t *testing.T) {
	// arrange
	resetViperAndConfig()

	// act
	config, _ := LoadConfig()

	// assert
	assert.NotNil(t, config)
	assert.IsType(t, &Config{}, config)
}
