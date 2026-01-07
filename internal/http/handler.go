package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexduzi/labcloudrun/internal/model"
)

type HttpHandler struct {
	Addr string
}

func NewHttpHandler(addr string) *HttpHandler {
	return &HttpHandler{
		Addr: addr,
	}
}

func (h *HttpHandler) GetTemperatureByCep(w http.ResponseWriter, r *http.Request) {
	temp := model.TemperatureResponse{
		Celsius:    28.5,
		Fahrenheit: 28.5,
		Kelvin:     30.0,
	}
	json.NewEncoder(w).Encode(temp)
}

func (h *HttpHandler) ListenAndServe() error {
	http.HandleFunc("/", h.GetTemperatureByCep)
	return http.ListenAndServe(fmt.Sprintf(":%s", h.Addr), nil)
}
