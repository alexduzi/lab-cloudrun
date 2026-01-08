package http

import (
	"net/http"

	"github.com/alexduzi/labcloudrun/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func (h HttpHandler) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandlerMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	router.GET("/readiness", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	router.GET("/:cep", h.GetTemperatureByCep)

	return router
}
