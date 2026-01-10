package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexduzi/labcloudrun/internal/client"
	cErrors "github.com/alexduzi/labcloudrun/internal/client/error"
	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/http/middleware"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func setupTestRouter(handler *HttpHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(middleware.ErrorHandlerMiddleware())

	r.GET("/:cep", handler.GetTemperatureByCep)

	return r
}

type HttpHandlerTestSuite struct {
	suite.Suite
	router            *gin.Engine
	cepClientStub     *client.CepClientStub
	weatherClientStub *client.WeatherClientStub
}

func (h *HttpHandlerTestSuite) SetupTest() {
	config := &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "test",
	}

	h.cepClientStub = client.NewCepClientStub(config)
	h.weatherClientStub = client.NewWeatherClientStub(config)

	handler := NewHttpHandler(config, h.cepClientStub, h.weatherClientStub)

	h.router = setupTestRouter(handler)
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_Success() {
	// arrange
	cep := "01001-000"
	city := "São Paulo"

	cepResponse := model.GetViacepResponseMock(cep)
	weatherResponse := model.GetWeatherResponseMock(city)

	expectedTempC := 32.2
	expectedTempF := 89.96  // 32.2 * 1.8 + 32
	expectedTempK := 305.35 // 32.2 + 273.15

	ctx := context.Background()

	h.cepClientStub.On("GetCep", ctx, cep).Return(cepResponse, nil)

	h.weatherClientStub.On("GetWeather", ctx, city).Return(weatherResponse, nil)

	// act
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/"+cep, nil)
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusOK, w.Code)

	var response model.TemperatureResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), expectedTempC, response.Celsius)
	assert.Equal(h.Suite.T(), expectedTempF, response.Fahrenheit)
	assert.Equal(h.Suite.T(), expectedTempK, response.Kelvin)

	h.cepClientStub.AssertExpectations(h.Suite.T())
	h.weatherClientStub.AssertExpectations(h.Suite.T())
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_CepNotFound() {
	// arrange
	cep := "11001-000"

	cepResponse := model.GetViacepResponseMock(cep)
	erro := "true"
	cepResponse.Erro = &erro

	ctx := context.Background()

	h.cepClientStub.On("GetCep", ctx, cep).Return(cepResponse, nil)

	// act
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/"+cep, nil)
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusNotFound, w.Code)

	var response model.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "can not find zipcode", response.Message)

	h.cepClientStub.AssertExpectations(h.Suite.T())
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_InvalidZipCode() {
	// arrange
	cep := "111111111111111111111111111"

	ctx := context.Background()

	// act
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/"+cep, nil)
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusUnprocessableEntity, w.Code)

	var response model.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "invalid zipcode", response.Message)
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_CepClientError() {
	// arrange
	cep := "01001-000"

	cepClientError := cErrors.NewCepClientHTTPError(500)

	ctx := context.Background()

	h.cepClientStub.On("GetCep", ctx, cep).Return(nil, cepClientError)

	// act
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/"+cep, nil)
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusInternalServerError, w.Code)

	var response model.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "internal server error", response.Message)

	h.cepClientStub.AssertExpectations(h.Suite.T())
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_WeatherClientError() {
	// arrange
	cep := "01001-000"
	city := "São Paulo"

	cepResponse := model.GetViacepResponseMock(cep)
	weatherClientError := cErrors.NewWeatherClientHTTPError(500)

	ctx := context.Background()

	h.cepClientStub.On("GetCep", ctx, cep).Return(cepResponse, nil)
	h.weatherClientStub.On("GetWeather", ctx, city).Return(nil, weatherClientError)

	// act
	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/"+cep, nil)
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusInternalServerError, w.Code)

	var response model.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "internal server error", response.Message)

	h.cepClientStub.AssertExpectations(h.Suite.T())
	h.weatherClientStub.AssertExpectations(h.Suite.T())
}

func (h *HttpHandlerTestSuite) TestHttpHandler_GetTemperatureByCep_CepParamNotExists() {
	// arrange
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	config := &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "test",
	}
	handler := NewHttpHandler(config, h.cepClientStub, h.weatherClientStub)

	router.Use(middleware.ErrorHandlerMiddleware())
	router.GET("/", handler.GetTemperatureByCep)

	req, _ := http.NewRequest("GET", "/", nil)

	// act
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusNotFound, w.Code)

	var response model.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)
	assert.Equal(h.Suite.T(), "can not find zipcode", response.Message)
}

func TestHttpHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HttpHandlerTestSuite))
}
