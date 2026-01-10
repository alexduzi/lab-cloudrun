package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alexduzi/labcloudrun/internal/client"
	"github.com/alexduzi/labcloudrun/internal/config"
	h "github.com/alexduzi/labcloudrun/internal/http"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	os.Setenv("WEATHER_API_KEY", os.Getenv("WEATHER_API_KEY"))
	if os.Getenv("WEATHER_API_KEY") == "" {
		os.Setenv("WEATHER_API_KEY", "6321a9495e9e46d5b1b230823260701")
	}

	cfg, _ := config.LoadConfig()

	cepApiApiClient := client.NewCepClient(cfg)
	weatherApiClient := client.NewWeatherClient(cfg)

	h := h.NewHttpHandler(cfg, cepApiApiClient, weatherApiClient)

	return h.SetupRouter()
}

func TestIntegration_GetTemperature_Success(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/temperature/01001-000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "temp_C")
	assert.Contains(t, w.Body.String(), "temp_F")
	assert.Contains(t, w.Body.String(), "temp_K")
}

func TestIntegration_GetTemperature_BlankCep(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	messageJson, _ := json.Marshal(model.ErrorResponse{
		Message: "can not find zipcode",
	})

	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/temperature/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, string(messageJson), w.Body.String())
}

func TestIntegration_GetTemperature_NonExistentCep(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	messageJson, _ := json.Marshal(model.ErrorResponse{
		Message: "can not find zipcode",
	})

	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/temperature/11001-000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, string(messageJson), w.Body.String())
}

func TestIntegration_GetTemperature_InvalidCep(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	messageJson, _ := json.Marshal(model.ErrorResponse{
		Message: "invalid zipcode",
	})

	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/temperature/11111111111111111111111111", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, string(messageJson), w.Body.String())
}
