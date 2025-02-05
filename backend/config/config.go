package config

import (
	"log"
	"os"
)

type AppConfig struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	RedirectURI         string
	FrontendURL         string
}

var Config AppConfig

func InitConfig() {
	Config = AppConfig{
		SpotifyClientID:     getEnvOrPanic("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: getEnvOrPanic("SPOTIFY_CLIENT_SECRET"),
		RedirectURI:         getEnvOrDefault("REDIRECT_URI", "http://localhost:8080/api/callback"),
		FrontendURL:         getEnvOrDefault("FRONTEND_URL", "http://localhost:3000"),
	}
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return value
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
