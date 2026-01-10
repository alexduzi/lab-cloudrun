package http

import (
	"log/slog"
	"net/http"

	"github.com/alexduzi/labcloudrun/internal/conversor"
	hErrors "github.com/alexduzi/labcloudrun/internal/http/error"
	"github.com/gin-gonic/gin"
)

// GetTemperatureWithoutCep handles requests without CEP parameter
func (h *HttpHandler) GetTemperatureWithoutCep(c *gin.Context) {
	slog.Error("CEP parameter not provided in request")
	_ = c.Error(hErrors.CepParamNotExists)
}

// GetTemperatureByCep godoc
// @Summary Get Temperature by CEP
// @Description Get temperature information by Brazilian postal code (CEP)
// @Tags weather
// @Accept json
// @Produce json
// @Param cep path string true "Brazilian postal code (CEP)" example(01310100)
// @Success 200 {object} model.TemperatureResponse "Temperature in Celsius, Fahrenheit and Kelvin"
// @Failure 404 {object} model.ErrorResponse "can not find zipcode"
// @Failure 422 {object} model.ErrorResponse "invalid zipcode"
// @Router /api/v1/temperature/{cep} [get]
func (h *HttpHandler) GetTemperatureByCep(c *gin.Context) {
	cep, _ := c.Params.Get("cep")

	if !h.cepRegex.MatchString(cep) {
		slog.Error("Invalid CEP format", "cep", cep)
		_ = c.Error(hErrors.CepInvalid)
		return
	}

	cepModel, err := h.cepApiClient.GetCep(c.Request.Context(), cep)
	if err != nil {
		slog.Error("Failed to get CEP information", "cep", cep, "error", err)
		_ = c.Error(err)
		return
	}

	if cepModel.Erro != nil {
		slog.Error("CEP not found", "cep", cep)
		_ = c.Error(hErrors.CepCantFind)
		return
	}

	weatherModel, err := h.weatherApiClient.GetWeather(c.Request.Context(), cepModel.Localidade)
	if err != nil {
		slog.Error("Failed to get weather information", "location", cepModel.Localidade, "cep", cep, "error", err)
		_ = c.Error(err)
		return
	}

	temp := conversor.ConvertWeatherResponse(*weatherModel)

	c.JSON(http.StatusOK, temp)
}
