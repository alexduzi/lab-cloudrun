package http

import (
	"github.com/alexduzi/labcloudrun/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func (h HttpHandler) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandlerMiddleware())

	router.GET("/:cep", h.GetTemperatureByCep)

	return router
}
