package http

import (
	"regexp"

	"github.com/alexduzi/labcloudrun/internal/client"
)

type HttpHandler struct {
	Addr             string
	cepApiClient     client.CepClientInterface
	weatherApiClient client.WeatherClientInterface
	cepRegex         *regexp.Regexp
}

func NewHttpHandler(
	addr string,
	cepApiClient client.CepClientInterface,
	weatherApiClient client.WeatherClientInterface) *HttpHandler {

	return &HttpHandler{
		Addr:             addr,
		cepApiClient:     cepApiClient,
		weatherApiClient: weatherApiClient,
		cepRegex:         regexp.MustCompile(`^\d{5}-?\d{3}$`),
	}
}
