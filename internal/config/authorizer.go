package config

import "os"

type Config struct {
	CardService           *CardService
	ExternalAuthorization *ExternalAuthorization
}

type CardService struct {
	Url string
}

type ExternalAuthorization struct {
	Url string
}

func ProvideConfig() *Config {
	return &Config{
		CardService:           &CardService{Url: os.Getenv("CARD_SERVICE_URL")},
		ExternalAuthorization: &ExternalAuthorization{Url: os.Getenv("EXTERNAL_AUTHORIZATION_URL")},
	}
}
