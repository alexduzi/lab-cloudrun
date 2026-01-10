package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WeatherClientStub struct {
	mock.Mock
	config *config.Config
	client *http.Client
}

func NewWeatherClientStub(cfg *config.Config) *WeatherClientStub {
	return &WeatherClientStub{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w *WeatherClientStub) GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error) {
	args := w.Called(ctx, city)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.WeatherResponse), nil
}

type WeatherClientStubTestSuite struct {
	suite.Suite
	config *config.Config
	client *WeatherClientStub
}

func (suite *WeatherClientStubTestSuite) SetupTest() {
	suite.config = &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "test",
	}
	suite.client = NewWeatherClientStub(suite.config)
}

func (suite *WeatherClientStubTestSuite) TestNewWeatherClientStub() {
	assert.NotNil(suite.T(), suite.client)
	assert.NotNil(suite.T(), suite.client.config)
	assert.NotNil(suite.T(), suite.client.client)
	assert.Equal(suite.T(), 10*time.Second, suite.client.client.Timeout)
}

func (suite *WeatherClientStubTestSuite) TestGetWeather_WithDifferentCities() {
	ctx := context.Background()

	testCases := []struct {
		name string
		city string
	}{
		{"Cidade brasileira", "São Paulo"},
		{"Cidade com espaços", "Rio de Janeiro"},
		{"Cidade internacional", "New York"},
		{"Cidade com caracteres especiais", "Belém"},
		{"Cidade vazia", ""},
	}

	resMock := &model.WeatherResponse{}

	for _, tc := range testCases {
		resMock.Location.Name = tc.city
		if tc.city != "" {
			suite.client.On("GetWeather", ctx, tc.city).Return(resMock, nil)
		} else {
			suite.client.On("GetWeather", ctx, tc.city).Return(nil, fmt.Errorf("Parameter q is missing."))
		}

		suite.Run(tc.name, func() {
			result, err := suite.client.GetWeather(ctx, tc.city)
			if tc.city != "" {
				assert.NotNil(suite.T(), result)
				assert.Nil(suite.T(), err)
			} else {
				assert.Nil(suite.T(), result)
				assert.NotNil(suite.T(), err)
				assert.Errorf(suite.T(), err, "Parameter q is missing.")
			}

		})
	}

	suite.client.AssertExpectations(suite.T())
}

func (suite *WeatherClientStubTestSuite) TestGetWeather_WithContext() {
	testCases := []struct {
		name string
		ctx  context.Context
	}{
		{"Context.Background", context.Background()},
		{"Context.TODO", context.TODO()},
	}

	resMock := &model.WeatherResponse{}

	for _, tc := range testCases {
		suite.client.On("GetWeather", tc.ctx, "São Paulo").Return(resMock, nil)

		suite.Run(tc.name, func() {
			result, err := suite.client.GetWeather(tc.ctx, "São Paulo")
			assert.NotNil(suite.T(), result)
			assert.Nil(suite.T(), err)
		})
	}

	suite.client.AssertExpectations(suite.T())
}

func (suite *WeatherClientStubTestSuite) TestGetWeather_ImplementsInterface() {
	var _ WeatherClientInterface = suite.client
}

func (suite *WeatherClientStubTestSuite) TestWeatherClientStub_HTTPClientConfiguration() {
	assert.NotNil(suite.T(), suite.client.client)
	assert.Equal(suite.T(), 10*time.Second, suite.client.client.Timeout)
}

func (suite *WeatherClientStubTestSuite) TestWeatherClientStub_ConfigInjection() {
	customConfig := &config.Config{
		Port:           "9090",
		WeatherAPIKey:  "custom-weather-key",
		ViaCEPBaseURL:  "https://custom.cep.com/",
		WeatherBaseURL: "https://custom.weather.api.com/v1/current.json",
		GinMode:        "release",
	}

	client := NewWeatherClientStub(customConfig)

	assert.NotNil(suite.T(), client)
	assert.Equal(suite.T(), customConfig, client.config)
	assert.Equal(suite.T(), "9090", client.config.Port)
	assert.Equal(suite.T(), "custom-weather-key", client.config.WeatherAPIKey)
	assert.Equal(suite.T(), "https://custom.weather.api.com/v1/current.json", client.config.WeatherBaseURL)
}

func (suite *WeatherClientStubTestSuite) TestWeatherClientStub_APIKeyConfiguration() {
	assert.NotNil(suite.T(), suite.client.config)
	assert.Equal(suite.T(), "test-api-key", suite.client.config.WeatherAPIKey)
}

func TestWeatherClientStubSuite(t *testing.T) {
	suite.Run(t, new(WeatherClientStubTestSuite))
}

// Testes independentes usando assert

func TestNewWeatherClientStub_WithNilConfig(t *testing.T) {
	client := NewWeatherClientStub(nil)

	assert.NotNil(t, client)
	assert.Nil(t, client.config)
	assert.NotNil(t, client.client)
}

func TestWeatherClientStub_GetWeather_MultipleCallsConsistency(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "test-key",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)
	ctx := context.Background()

	resMock := &model.WeatherResponse{}
	resMock.Location.Name = "São Paulo"

	client.On("GetWeather", ctx, "São Paulo").Return(resMock, nil)

	for i := 0; i < 5; i++ {
		result, err := client.GetWeather(ctx, "São Paulo")
		assert.NotNil(t, result)
		assert.Nil(t, err)
	}

	client.AssertExpectations(t)
}

func TestWeatherClientStub_GetWeather_WithCanceledContext(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "test-key",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancela o contexto antes da chamada

	client.On("GetWeather", ctx, mock.Anything).Return(nil, context.Canceled)

	result, err := client.GetWeather(ctx, "São Paulo")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, context.Canceled)

	client.AssertExpectations(t)
}

func TestWeatherClientStub_GetWeather_WithTimeout(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "test-key",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(2 * time.Millisecond) // Garante que o timeout aconteceu

	client.On("GetWeather", ctx, mock.Anything).Return(nil, context.DeadlineExceeded)

	result, err := client.GetWeather(ctx, "São Paulo")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, context.DeadlineExceeded)

	client.AssertExpectations(t)
}

func TestWeatherClientStub_GetWeather_WithEmptyAPIKey(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)
	ctx := context.Background()

	client.On("GetWeather", ctx, mock.Anything).Return(nil, fmt.Errorf("API key is invalid."))

	result, err := client.GetWeather(ctx, "São Paulo")

	assert.Nil(t, result)
	assert.Errorf(t, err, "API key is invalid.")

	client.AssertExpectations(t)
}

func TestWeatherClientStub_GetWeather_WithSpecialCharacters(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "test-key",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)
	ctx := context.Background()

	testCases := []string{
		"São Paulo",
		"Belém",
		"Brasília",
		"Goiânia",
		"Florianópolis",
	}

	resMock := &model.WeatherResponse{}

	for _, city := range testCases {
		resMock.Location.Name = city
		client.On("GetWeather", ctx, city).Return(resMock, nil)

		t.Run(city, func(t *testing.T) {
			result, err := client.GetWeather(ctx, city)
			assert.NotNil(t, result)
			assert.Nil(t, err)
		})
	}

	client.AssertExpectations(t)
}

func TestWeatherClientStub_HTTPClient_Timeout(t *testing.T) {
	cfg := &config.Config{
		WeatherAPIKey:  "test-key",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
	}
	client := NewWeatherClientStub(cfg)

	assert.Equal(t, 10*time.Second, client.client.Timeout)
}

func TestWeatherClientStub_ConfigurationValues(t *testing.T) {
	expectedAPIKey := "my-secret-api-key"
	expectedBaseURL := "https://custom.weatherapi.com/v1/current.json"

	cfg := &config.Config{
		WeatherAPIKey:  expectedAPIKey,
		WeatherBaseURL: expectedBaseURL,
		Port:           "3000",
		GinMode:        "release",
	}

	client := NewWeatherClientStub(cfg)

	assert.Equal(t, expectedAPIKey, client.config.WeatherAPIKey)
	assert.Equal(t, expectedBaseURL, client.config.WeatherBaseURL)
	assert.Equal(t, "3000", client.config.Port)
	assert.Equal(t, "release", client.config.GinMode)
}
