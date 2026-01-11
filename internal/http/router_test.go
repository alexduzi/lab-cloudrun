package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexduzi/labcloudrun/internal/client"
	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	handler       *HttpHandler
	router        *gin.Engine
	cepClient     *client.CepClientStub
	weatherClient *client.WeatherClientStub
	config        *config.Config
}

func (s *RouterTestSuite) SetupTest() {
	s.config = &config.Config{
		Port:           "8080",
		WeatherAPIKey:  "test-api-key",
		ViaCEPBaseURL:  "https://viacep.com.br/ws/{cep}/json/",
		WeatherBaseURL: "http://api.weatherapi.com/v1/current.json",
		GinMode:        gin.TestMode,
	}

	s.cepClient = client.NewCepClientStub(s.config)
	s.weatherClient = client.NewWeatherClientStub(s.config)

	s.handler = NewHttpHandler(s.config, s.cepClient, s.weatherClient)
	s.router = s.handler.SetupRouter()
}

func (s *RouterTestSuite) TestSetupRouter_GinModeSet() {
	assert.Equal(s.T(), gin.TestMode, gin.Mode())
}

func (s *RouterTestSuite) TestSetupRouter_ReturnsRouter() {
	assert.NotNil(s.T(), s.router)
}

func (s *RouterTestSuite) TestSetupRouter_HealthEndpointRegistered() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	s.router.ServeHTTP(w, req)

	assert.NotEqual(s.T(), http.StatusNotFound, w.Code)
	assert.Equal(s.T(), http.StatusOK, w.Code)
}

func (s *RouterTestSuite) TestSetupRouter_ReadinessEndpointRegistered() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/readiness", nil)
	s.router.ServeHTTP(w, req)

	assert.NotEqual(s.T(), http.StatusNotFound, w.Code)
	assert.Equal(s.T(), http.StatusOK, w.Code)
}

func (s *RouterTestSuite) TestSetupRouter_SwaggerEndpointRegistered() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/swagger/doc.json", nil)
	s.router.ServeHTTP(w, req)

	routes := s.router.Routes()
	swaggerExists := false
	for _, route := range routes {
		if route.Path == "/swagger/*any" {
			swaggerExists = true
			break
		}
	}
	assert.True(s.T(), swaggerExists, "Swagger route should be registered")
}

func (s *RouterTestSuite) TestSetupRouter_TemperatureWithoutCepEndpointRegistered() {
	routes := s.router.Routes()
	tempWithoutCepExists := false
	for _, route := range routes {
		if route.Path == "/api/v1/temperature/" {
			tempWithoutCepExists = true
			break
		}
	}
	assert.True(s.T(), tempWithoutCepExists, "Temperature without CEP route should be registered")
}

func (s *RouterTestSuite) TestSetupRouter_APIGroupExists() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/nonexistent", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

func (s *RouterTestSuite) TestSetupRouter_RouteCount() {
	routes := s.router.Routes()

	// 1. /health
	// 2. /readiness
	// 3. /swagger/*any
	// 4. /api/v1/temperature/
	// 5. /api/v1/temperature/:cep
	assert.GreaterOrEqual(s.T(), len(routes), 5, "Should have at least 5 routes registered")
}

func (s *RouterTestSuite) TestSetupRouter_ErrorMiddlewareApplied() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/temperature/invalid-cep", nil)
	s.router.ServeHTTP(w, req)

	assert.NotEqual(s.T(), http.StatusInternalServerError, w.Code)
}

func (s *RouterTestSuite) TestSetupRouter_NonExistentRoute() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

func (s *RouterTestSuite) TestSetupRouter_MethodNotAllowed() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/health", nil)
	s.router.ServeHTTP(w, req)

	assert.True(s.T(), w.Code == http.StatusNotFound || w.Code == http.StatusMethodNotAllowed)
}

func (s *RouterTestSuite) TestSetupRouter_WithDifferentGinModes() {
	modes := []string{gin.TestMode, gin.ReleaseMode, gin.DebugMode}

	for _, mode := range modes {
		s.config.GinMode = mode
		handler := NewHttpHandler(s.config, s.cepClient, s.weatherClient)
		router := handler.SetupRouter()

		assert.NotNil(s.T(), router)
		assert.Equal(s.T(), mode, gin.Mode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)
		assert.Equal(s.T(), http.StatusOK, w.Code)
	}
}

func (s *RouterTestSuite) TestSetupRouter_AllEndpointsTableDriven() {
	tests := []struct {
		name      string
		routePath string
	}{
		{
			name:      "Health endpoint",
			routePath: "/health",
		},
		{
			name:      "Readiness endpoint",
			routePath: "/readiness",
		},
		{
			name:      "Swagger documentation",
			routePath: "/swagger/*any",
		},
		{
			name:      "Temperature by CEP",
			routePath: "/api/v1/temperature/:cep",
		},
		{
			name:      "Temperature without CEP",
			routePath: "/api/v1/temperature/",
		},
	}

	routes := s.router.Routes()
	registeredPaths := make(map[string]bool)
	for _, route := range routes {
		registeredPaths[route.Path] = true
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			assert.True(s.T(), registeredPaths[tt.routePath], "Route %s should be registered", tt.routePath)
		})
	}
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}
