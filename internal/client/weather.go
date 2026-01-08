package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alexduzi/labcloudrun/internal/model"
)

type WeatherClientInterface interface {
	GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error)
}

type WeatherClient struct {
	appKey     string
	baseApiUrl string
	client     *http.Client
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		appKey:     "6321a9495e9e46d5b1b230823260701",
		baseApiUrl: "http://api.weatherapi.com/v1/current.json?key={appKey}&q={city}&aqi=no",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w WeatherClient) GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error) {
	weatherApiUrl := strings.Replace(w.baseApiUrl, "{appKey}", w.appKey, 1)
	weatherApiUrl = strings.Replace(weatherApiUrl, "{city}", url.QueryEscape(city), 1)

	req, err := http.NewRequestWithContext(ctx, "GET", weatherApiUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
