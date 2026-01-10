package http

import (
	_ "github.com/alexduzi/labcloudrun/docs"
	"github.com/alexduzi/labcloudrun/internal/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h HttpHandler) SetupRouter() *gin.Engine {
	// Set Gin mode based on configuration
	gin.SetMode(h.config.GinMode)

	router := gin.Default()

	router.Use(middleware.ErrorHandlerMiddleware())

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health and Readiness endpoints
	router.GET("/health", h.HealthCheck)
	router.GET("/readiness", h.ReadinessCheck)

	// Weather endpoint
	v1 := router.Group("/api/v1")
	v1.GET("/temperature/", h.GetTemperatureWithoutCep)
	v1.GET("/temperature/:cep", h.GetTemperatureByCep)

	return router
}
