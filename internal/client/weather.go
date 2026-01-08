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
	GetWeather(ctx context.Context, city string) (*model.WheaterResponse, error)
}

type WeatherClient struct {
	appKey     string
	baseApiUrl string
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		appKey:     "6321a9495e9e46d5b1b230823260701",
		baseApiUrl: "http://api.weatherapi.com/v1/current.json?key={appKey}&q={city}&aqi=no",
	}
}

func (w WeatherClient) GetWeather(ctx context.Context, city string) (*model.WheaterResponse, error) {
	client := http.Client{}

	w.baseApiUrl = strings.Replace(w.baseApiUrl, "{appKey}", w.appKey, 1)
	w.baseApiUrl = strings.Replace(w.baseApiUrl, "{city}", city, 1)

	req, err := http.NewRequestWithContext(ctx, "GET", w.baseApiUrl, nil)
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

	var wheatherRes model.WheaterResponse
	err = json.Unmarshal(body, &wheatherRes)
	if err != nil {
		return nil, err
	}

	return &wheatherRes, nil
}
