package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

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
		client:     &http.Client{},
	}
}

func (w WeatherClient) GetWeather(ctx context.Context, city string) (*model.WeatherResponse, error) {
	weatherApiUrl := strings.Replace(w.baseApiUrl, "{appKey}", w.appKey, 1)
	weatherApiUrl = strings.Replace(weatherApiUrl, "{city}", city, 1)

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

	var wheatherRes model.WeatherResponse
	err = json.Unmarshal(body, &wheatherRes)
	if err != nil {
		return nil, err
	}

	return &wheatherRes, nil
}
