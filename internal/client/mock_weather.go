package client

import (
	"context"
	"net/http"
	"time"

	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
	"github.com/stretchr/testify/mock"
)

type WeatherClientStub struct {
	mock.Mock
	config *config.Config
	client *http.Client
}

func NewWeatherClientStub(cfg *config.Config) *WeatherClientStub {
	return &WeatherClientStub{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w *WeatherClientStub) GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error) {
	args := w.Called(ctx, city)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.WeatherResponse), nil
}
