package http

import (
	"regexp"

	"github.com/alexduzi/labcloudrun/internal/client"
)

type HttpHandler struct {
	Addr              string
	cepApiClient      client.CepClientInterface
	wheatherApiClient client.WeatherClientInterface
	cepRegex          *regexp.Regexp
}

func NewHttpHandler(
	addr string,
	cepApiClient client.CepClientInterface,
	wheatherApiClient client.WeatherClientInterface) *HttpHandler {

	return &HttpHandler{
		Addr:              addr,
		cepApiClient:      cepApiClient,
		wheatherApiClient: wheatherApiClient,
		cepRegex:          regexp.MustCompile(`^\d{5}-?\d{3}$`),
	}
}
