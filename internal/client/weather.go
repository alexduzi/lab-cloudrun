package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	cErrors "github.com/alexduzi/labcloudrun/internal/client/error"
	"github.com/alexduzi/labcloudrun/internal/config"
	"github.com/alexduzi/labcloudrun/internal/model"
)

type WeatherClientInterface interface {
	GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error)
}

type WeatherClient struct {
	config *config.Config
	client *http.Client
}

func NewWeatherClient(cfg *config.Config) *WeatherClient {
	return &WeatherClient{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w WeatherClient) GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error) {
	weatherApiUrl := fmt.Sprintf("%s?key=%s&q=%s&aqi=no",
		w.config.WeatherBaseURL,
		w.config.WeatherAPIKey,
		url.QueryEscape(city))

	req, err := http.NewRequestWithContext(ctx, "GET", weatherApiUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, cErrors.NewWeatherClientHTTPError(resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherRes model.WeatherResponse
	err = json.Unmarshal(body, &weatherRes)
	if err != nil {
		return nil, err
	}

	return &weatherRes, nil
}
