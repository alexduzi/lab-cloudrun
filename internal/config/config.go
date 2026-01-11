package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	WeatherAPIKey  string
	ViaCEPBaseURL  string
	WeatherBaseURL string
	GinMode        string
}

var AppConfig *Config

// LoadConfig loads configuration from environment variables using Viper
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set default values
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("PORT")
	}
	if port == "" {
		port = "8080" // Default fallback
	}

	viper.SetDefault("VIA_CEP_BASE_URL", "https://viacep.com.br/ws/{cep}/json/")
	viper.SetDefault("WEATHER_BASE_URL", "http://api.weatherapi.com/v1/current.json")
	viper.SetDefault("GIN_MODE", "debug") // debug, release, or test

	// Try to read .env file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found, using environment variables and defaults")
		} else {
			log.Printf("Error reading config file: %v", err)
		}
	}

	config := &Config{
		Port:           port,
		WeatherAPIKey:  viper.GetString("WEATHER_API_KEY"),
		ViaCEPBaseURL:  viper.GetString("VIA_CEP_BASE_URL"),
		WeatherBaseURL: viper.GetString("WEATHER_BASE_URL"),
		GinMode:        viper.GetString("GIN_MODE"),
	}

	// Validate required fields
	if config.WeatherAPIKey == "" {
		log.Println("Warning: WEATHER_API_KEY is not set")
	}

	AppConfig = config
	return config, nil
}

// GetConfig returns the current configuration
func GetConfig() *Config {
	if AppConfig == nil {
		log.Fatal("Config not initialized. Call LoadConfig() first.")
	}
	return AppConfig
}
