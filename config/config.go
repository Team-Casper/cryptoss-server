package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBPath        string
	ListeningAddr string
	SMSEndpoint   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	return &Config{
		DBPath:        os.Getenv("DBPATH"),
		ListeningAddr: os.Getenv("LISTENING_ADDR"),
		SMSEndpoint:   os.Getenv("SMS_ENDPOINT"),
	}, nil
}
