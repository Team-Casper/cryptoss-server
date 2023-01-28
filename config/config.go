package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
)

type Config struct {
	DBPath        string
	ListeningAddr string
	FromNumber    string
	SMSEndpoint   string
	AccessKeyId   string
	SecretKey     string
}

func LoadConfig() (*Config, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath := path.Join(pwd, ".env")

	if err = godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("failed to load .env file(%s): %w", envPath, err)
	}

	return &Config{
		DBPath:        os.Getenv("DBPATH"),
		ListeningAddr: os.Getenv("LISTENING_ADDR"),
		FromNumber:    os.Getenv("FROM_NUMBER"),
		SMSEndpoint:   os.Getenv("SMS_ENDPOINT"),
		AccessKeyId:   os.Getenv("NCP_ACCESS_KEY_ID"),
		SecretKey:     os.Getenv("NCP_SECRET_KEY"),
	}, nil
}
