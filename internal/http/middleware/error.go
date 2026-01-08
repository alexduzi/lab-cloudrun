package middleware

import (
	"errors"
	"net/http"

	hErrors "github.com/alexduzi/labcloudrun/internal/http/error"
	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if errors.Is(err, hErrors.CepParamNotExists) {
				c.JSON(http.StatusNotFound, "can not find zipcode")
				return
			}

			if errors.Is(err, hErrors.CepCantFind) {
				c.JSON(http.StatusNotFound, "can not find zipcode")
				return
			}

			if errors.Is(err, hErrors.CepInvalid) {
				c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
				return
			}
		}
	}
}
