package http

import (
	"net/http"

	"github.com/alexduzi/labcloudrun/internal/conversor"
	hErrors "github.com/alexduzi/labcloudrun/internal/http/error"
	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) GetTemperatureByCep(c *gin.Context) {
	cep, exists := c.Params.Get("cep")
	if !exists {
		c.Error(hErrors.CepParamNotExists)
		return
	}

	if !h.cepRegex.MatchString(cep) {
		c.Error(hErrors.CepInvalid)
		return
	}

	cepModel, err := h.cepApiClient.GetCep(c.Request.Context(), cep)
	if err != nil {
		c.Error(err)
		return
	}

	if cepModel.Erro != nil {
		c.Error(hErrors.CepCantFind)
		return
	}

	weatherModel, err := h.wheatherApiClient.GetWeather(c.Request.Context(), cepModel.Localidade)
	if err != nil {
		c.Error(err)
		return
	}

	temp := conversor.ConvertWeatherResponse(*weatherModel)

	c.JSON(http.StatusOK, temp)
}
