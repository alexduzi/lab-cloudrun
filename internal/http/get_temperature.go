package http

import (
	"net/http"

	"github.com/alexduzi/labcloudrun/internal/conversor"
	hErrors "github.com/alexduzi/labcloudrun/internal/http/error"
	"github.com/gin-gonic/gin"
)

// GetTemperatureByCep godoc
// @Summary Get Weather
// @Description Get Weather by zipcode
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} model.WeatherResponse
// @Failure 500 {string} string
// @Router / [get]
func (h *HttpHandler) GetTemperatureByCep(c *gin.Context) {
	cep, exists := c.Params.Get("cep")
	if !exists {
		_ = c.Error(hErrors.CepParamNotExists)
		return
	}

	if !h.cepRegex.MatchString(cep) {
		_ = c.Error(hErrors.CepInvalid)
		return
	}

	cepModel, err := h.cepApiClient.GetCep(c.Request.Context(), cep)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if cepModel.Erro != nil {
		_ = c.Error(hErrors.CepCantFind)
		return
	}

	weatherModel, err := h.weatherApiClient.GetWeather(c.Request.Context(), cepModel.Localidade)
	if err != nil {
		_ = c.Error(err)
		return
	}

	temp := conversor.ConvertWeatherResponse(*weatherModel)

	c.JSON(http.StatusOK, temp)
}
