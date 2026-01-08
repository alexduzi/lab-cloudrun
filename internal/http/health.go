package http

import (
	"net/http"
	"time"

	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Health Check
// @Description Check if the service is healthy and running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} model.StatusResponse
// @Router /health [get]
func (h *HttpHandler) HealthCheck(c *gin.Context) {
	response := model.StatusResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "lab-cloudrun-api",
	}

	c.JSON(http.StatusOK, response)
}

// ReadinessCheck godoc
// @Summary Readiness Check
// @Description Check if the service is ready to accept traffic
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} model.StatusResponse
// @Router /readiness [get]
func (h *HttpHandler) ReadinessCheck(c *gin.Context) {
	response := model.StatusResponse{
		Status:    "ready",
		Timestamp: time.Now(),
		Service:   "lab-cloudrun-api",
	}

	c.JSON(http.StatusOK, response)
}
