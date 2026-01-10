package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	cErrors "github.com/alexduzi/labcloudrun/internal/client/error"
	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
)

type CepClientInterface interface {
	GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error)
}

type CepClient struct {
	config *config.Config
	client *http.Client
}

func NewCepClient(cfg *config.Config) *CepClient {
	return &CepClient{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c CepClient) GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error) {
	cepApiUrl := strings.Replace(c.config.ViaCEPBaseURL, "{cep}", cep, 1)

	req, err := http.NewRequestWithContext(ctx, "GET", cepApiUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, cErrors.NewCepClientHTTPError(resp.StatusCode)
	}

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
