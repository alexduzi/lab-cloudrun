package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CepClientStub struct {
	config *config.Config
	client *http.Client
}

func NewCepClientStub(cfg *config.Config) *CepClientStub {
	return &CepClientStub{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c CepClientStub) GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error) {
	return nil, nil
}

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

	result, err := suite.client.GetCep(ctx, "01310100")

	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), err)
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

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			result, err := suite.client.GetCep(ctx, tc.cep)
			assert.Nil(suite.T(), result)
			assert.Nil(suite.T(), err)
		})
	}
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

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			result, err := suite.client.GetCep(tc.ctx, "01310100")
			assert.Nil(suite.T(), result)
			assert.Nil(suite.T(), err)
		})
	}
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

	// Múltiplas chamadas devem retornar o mesmo resultado
	for i := 0; i < 5; i++ {
		result, err := client.GetCep(ctx, "01310100")
		assert.Nil(t, result)
		assert.Nil(t, err)
	}
}

func TestCepClientStub_GetCep_WithCanceledContext(t *testing.T) {
	cfg := &config.Config{
		ViaCEPBaseURL: "https://viacep.com.br/ws/{cep}/json/",
	}
	client := NewCepClientStub(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancela o contexto antes da chamada

	result, err := client.GetCep(ctx, "01310100")

	// O stub não verifica contexto cancelado, retorna nil/nil sempre
	assert.Nil(t, result)
	assert.Nil(t, err)
}
