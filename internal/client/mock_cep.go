package client

import (
	"context"
	"net/http"
	"time"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/mock"
)

type CepClientStub struct {
	mock.Mock
	config *config.Config
	client *http.Client
}

func NewCepClientStub(cfg *config.Config) *CepClientStub {
	return &CepClientStub{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *CepClientStub) GetCep(ctx context.Context, cep string) (*model.ViacepResponse, error) {
	args := c.Called(ctx, cep)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ViacepResponse), nil
}
