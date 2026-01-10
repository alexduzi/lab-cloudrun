package middleware

import (
	"errors"
	"net/http"

	cErrors "github.com/alexduzi/labcloudrun/internal/client/error"
	hErrors "github.com/alexduzi/labcloudrun/internal/http/error"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle errors if response hasn't been written yet
		if len(c.Errors) > 0 && !c.Writer.Written() {
			err := c.Errors.Last().Err

			if errors.Is(err, hErrors.CepParamNotExists) {
				c.JSON(http.StatusNotFound, model.ErrorResponse{
					Message: "can not find zipcode",
				})
				return
			}

			if errors.Is(err, hErrors.CepCantFind) {
				c.JSON(http.StatusNotFound, model.ErrorResponse{
					Message: "can not find zipcode",
				})
				return
			}

			if errors.Is(err, hErrors.CepInvalid) {
				c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
					Message: "invalid zipcode",
				})
				return
			}

			// Handle CEP client errors
			if errors.Is(err, cErrors.CepClientBadRequest) ||
				errors.Is(err, cErrors.CepClientNotFound) ||
				errors.Is(err, cErrors.CepClientInternalError) ||
				errors.Is(err, cErrors.CepClientUnexpectedError) {
				c.JSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: "internal server error",
				})
				return
			}

			// Handle Weather client errors
			if errors.Is(err, cErrors.WeatherClientBadRequest) ||
				errors.Is(err, cErrors.WeatherClientNotFound) ||
				errors.Is(err, cErrors.WeatherClientInternalError) ||
				errors.Is(err, cErrors.WeatherClientUnexpectedError) {
				c.JSON(http.StatusInternalServerError, model.ErrorResponse{
					Message: "internal server error",
				})
				return
			}

			// Handle unexpected errors
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "internal server error",
			})
		}
	}
}
