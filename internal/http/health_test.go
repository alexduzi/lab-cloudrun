package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HealthHandlerTestSuite struct {
	suite.Suite
	router  *gin.Engine
	handler *HttpHandler
}

func (h *HealthHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        "test",
	}

	h.handler = NewHttpHandler(cfg, nil, nil)
	h.router = gin.New()
}

func (h *HealthHandlerTestSuite) TestHealthCheck_Success() {
	// arrange
	h.router.GET("/health", h.handler.HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// act
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusOK, w.Code)

	var response model.StatusResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "healthy", response.Status)
	assert.Equal(h.Suite.T(), "lab-cloudrun-api", response.Service)
	assert.NotZero(h.Suite.T(), response.Timestamp)
}

func (h *HealthHandlerTestSuite) TestReadinessCheck_Success() {
	// arrange
	h.router.GET("/readiness", h.handler.ReadinessCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/readiness", nil)

	// act
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusOK, w.Code)

	var response model.StatusResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(h.Suite.T(), err)

	assert.Equal(h.Suite.T(), "ready", response.Status)
	assert.Equal(h.Suite.T(), "lab-cloudrun-api", response.Service)
	assert.NotZero(h.Suite.T(), response.Timestamp)
}

func (h *HealthHandlerTestSuite) TestHealthCheck_ResponseContentType() {
	// arrange
	h.router.GET("/health", h.handler.HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// act
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusOK, w.Code)
	assert.Contains(h.Suite.T(), w.Header().Get("Content-Type"), "application/json")
}

func (h *HealthHandlerTestSuite) TestReadinessCheck_ResponseContentType() {
	// arrange
	h.router.GET("/readiness", h.handler.ReadinessCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/readiness", nil)

	// act
	h.router.ServeHTTP(w, req)

	// assert
	assert.Equal(h.Suite.T(), http.StatusOK, w.Code)
	assert.Contains(h.Suite.T(), w.Header().Get("Content-Type"), "application/json")
}

func TestHealthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthHandlerTestSuite))
}
