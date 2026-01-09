package http

import (
	"regexp"

	"github.com/alexduzi/labcloudrun/internal/client"
	"github.com/alexduzi/labcloudrun/internal/config"
)

type HttpHandler struct {
	config           *config.Config
	cepApiClient     client.CepClientInterface
	weatherApiClient client.WeatherClientInterface
	cepRegex         *regexp.Regexp
}

func NewHttpHandler(
	cfg *config.Config,
	cepApiClient client.CepClientInterface,
	weatherApiClient client.WeatherClientInterface) *HttpHandler {

	return &HttpHandler{
		config:           cfg,
		cepApiClient:     cepApiClient,
		weatherApiClient: weatherApiClient,
		cepRegex:         regexp.MustCompile(`^\d{5}-?\d{3}$`),
	}
}
