package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alexduzi/labcloudrun/internal/model"
)

type CepClientInterface interface {
	GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error)
}

type CepClient struct {
	baseCepUrl string
}

func NewCepClient() *CepClient {
	return &CepClient{
		baseCepUrl: "https://viacep.com.br/ws/{cep}/json/",
	}
}

func (c CepClient) GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", strings.Replace(c.baseCepUrl, "{cep}", cep, 1), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cepRes model.ViacepResponse
	err = json.Unmarshal(body, &cepRes)
	if err != nil {
		return nil, err
	}

	return &cepRes, nil
}
