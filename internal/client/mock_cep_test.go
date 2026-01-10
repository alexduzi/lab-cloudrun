package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// CepClientStubTestSuite é a test suite para CepClientStub
type CepClientStubTestSuite struct {
	suite.Suite
	config *config.Config
	client *CepClientStub
}

// SetupTest é executado antes de cada teste
func (suite *CepClientStubTestSuite) SetupTest() {
	suite.config = &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "test",
	}
	suite.client = NewCepClientStub(suite.config)
}

// TestNewCepClientStub testa a criação de um novo CepClientStub
func (suite *CepClientStubTestSuite) TestNewCepClientStub() {
	assert.NotNil(suite.T(), suite.client)
	assert.NotNil(suite.T(), suite.client.config)
	assert.NotNil(suite.T(), suite.client.client)
	assert.Equal(suite.T(), 10*time.Second, suite.client.client.Timeout)
}

// TestGetCep_ReturnsNil testa que GetCep retorna nil (comportamento do stub)
func (suite *CepClientStubTestSuite) TestGetCep_ReturnsNil() {
	ctx := context.Background()

	cep := "01310100"

	resMock := model.GetViacepResponseMock(cep)

	suite.client.On("GetCep", ctx, cep).Return(resMock, nil)

	result, err := suite.client.GetCep(ctx, cep)

	assert.NotNil(suite.T(), result)
	assert.Nil(suite.T(), err)

	suite.client.AssertExpectations(suite.T())
}

// TestGetCep_WithDifferentCeps testa GetCep com diferentes CEPs
func (suite *CepClientStubTestSuite) TestGetCep_WithDifferentCeps() {
	ctx := context.Background()

	testCases := []struct {
		name string
		cep  string
	}{
		{"CEP válido", "01310100"},
		{"CEP com hífen", "01310-100"},
		{"CEP inválido", "00000000"},
		{"CEP vazio", ""},
	}

	resMock := model.GetViacepResponseMock("cep")

	for _, tc := range testCases {
		resMock.Cep = tc.cep

		if tc.cep != "" {
			suite.client.On("GetCep", ctx, tc.cep).Return(resMock, nil)
		} else {
			suite.client.On("GetCep", ctx, tc.cep).Return(resMock, fmt.Errorf("Parameter cep is missing."))
		}

		suite.Run(tc.name, func() {

			result, err := suite.client.GetCep(ctx, tc.cep)
			if tc.cep != "" {
				assert.NotNil(suite.T(), result)
				assert.Nil(suite.T(), err)
			} else {
				assert.Nil(suite.T(), result)
				assert.NotNil(suite.T(), err)
				assert.Errorf(suite.T(), err, "Parameter cep is missing.")
			}
		})
	}

	suite.client.AssertExpectations(suite.T())
}

// TestGetCep_WithContext testa GetCep com diferentes contextos
func (suite *CepClientStubTestSuite) TestGetCep_WithContext() {
	testCases := []struct {
		name string
		ctx  context.Context
	}{
		{"Context.Background", context.Background()},
		{"Context.TODO", context.TODO()},
	}

	cep := "01310100"

	resMock := model.GetViacepResponseMock(cep)

	for _, tc := range testCases {
		suite.client.On("GetCep", tc.ctx, cep).Return(resMock, nil)

		suite.Run(tc.name, func() {
			result, err := suite.client.GetCep(tc.ctx, cep)
			assert.NotNil(suite.T(), result)
			assert.Nil(suite.T(), err)
		})
	}

	suite.client.AssertExpectations(suite.T())
}

// TestGetCep_ImplementsInterface verifica se CepClientStub implementa CepClientInterface
func (suite *CepClientStubTestSuite) TestGetCep_ImplementsInterface() {
	var _ CepClientInterface = suite.client
}

// TestCepClientStub_HTTPClientConfiguration testa a configuração do HTTP client
func (suite *CepClientStubTestSuite) TestCepClientStub_HTTPClientConfiguration() {
	assert.NotNil(suite.T(), suite.client.client)
	assert.Equal(suite.T(), 10*time.Second, suite.client.client.Timeout)
}

// TestCepClientStub_ConfigInjection testa a injeção de configuração
func (suite *CepClientStubTestSuite) TestCepClientStub_ConfigInjection() {
	customConfig := &config.Config{
		Port:           "9090",
		WeatherAPIKey:  "custom-key",
		ViaCEPBaseURL:  "https://custom.api.com/{cep}/",
		WeatherBaseURL: "https://custom.weather.com/",
		GinMode:        "release",
	}

	client := NewCepClientStub(customConfig)

	assert.NotNil(suite.T(), client)
	assert.Equal(suite.T(), customConfig, client.config)
	assert.Equal(suite.T(), "9090", client.config.Port)
	assert.Equal(suite.T(), "custom-key", client.config.WeatherAPIKey)
}

// TestCepClientStubSuite executa a test suite
func TestCepClientStubSuite(t *testing.T) {
	suite.Run(t, new(CepClientStubTestSuite))
}

// Testes independentes usando assert

func TestNewCepClientStub_WithNilConfig(t *testing.T) {
	client := NewCepClientStub(nil)

	assert.NotNil(t, client)
	assert.Nil(t, client.config)
	assert.NotNil(t, client.client)
}

func TestCepClientStub_GetCep_MultipleCallsConsistency(t *testing.T) {
	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
	}
	client := NewCepClientStub(cfg)
	ctx := context.Background()

	cep := "01310100"

	resMock := model.GetViacepResponseMock(cep)

	client.On("GetCep", ctx, cep).Return(resMock, nil)

	// Múltiplas chamadas devem retornar o mesmo resultado
	for i := 0; i < 5; i++ {
		result, err := client.GetCep(ctx, cep)
		assert.NotNil(t, result)
		assert.Nil(t, err)
	}

	client.AssertExpectations(t)
}

func TestCepClientStub_GetCep_WithCanceledContext(t *testing.T) {
	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
	}
	client := NewCepClientStub(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancela o contexto antes da chamada

	client.On("GetCep", ctx, mock.Anything).Return(nil, context.Canceled)

	result, err := client.GetCep(ctx, "01310100")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, context.Canceled)

	client.AssertExpectations(t)
}

func TestCepClientStub_GetCep_WithTimeout(t *testing.T) {
	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
	}
	client := NewCepClientStub(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(2 * time.Millisecond) // Garante que o timeout aconteceu

	client.On("GetCep", ctx, mock.Anything).Return(nil, context.DeadlineExceeded)

	result, err := client.GetCep(ctx, "01310100")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, context.DeadlineExceeded)

	client.AssertExpectations(t)
}

func TestCepClientStub_HTTPClient_Timeout(t *testing.T) {
	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
	}
	client := NewCepClientStub(cfg)

	assert.Equal(t, 10*time.Second, client.client.Timeout)
}

func TestCepClientStub_ConfigurationValues(t *testing.T) {
	expectedBaseURL := "https://viacep.com.br/ws/{cep}/json/"

	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
		Port:          "3000",
		GinMode:       "release",
	}
	client := NewCepClientStub(cfg)

	assert.Equal(t, expectedBaseURL, client.config.ViaCEPBaseURL)
	assert.Equal(t, "3000", client.config.Port)
	assert.Equal(t, "release", client.config.GinMode)
}
