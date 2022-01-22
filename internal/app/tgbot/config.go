package tgbot

import (
	"log"
	"os"
)

const (
	logLevelDebug = "debug"
)

type Config struct {
	GoogleApplicationCredentials string `toml:"google_application_credentials"`
	GoogleCloudProjectID         string `toml:"google_cloud_project_id"`
	LogLevel                     string `toml:"log_level"`
	TelegramApiToken             string `toml:"telegram_api_token"`
}

func NewConfig() *Config {
	return &Config{
		LogLevel: logLevelDebug,
	}
}

func (c *Config) LoadFromEnv() error {
	if len(c.GoogleApplicationCredentials) == 0 {
		if len(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")) == 0 {
			log.Fatal("GoogleApplicationCredentials config not set!")
		}
		c.GoogleApplicationCredentials = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}
	if len(c.GoogleCloudProjectID) == 0 {
		if len(os.Getenv("GOOGLE_CLOUD_PROJECT_ID")) == 0 {
			log.Fatal("GoogleCloudProjectID config not set!")
		}
		c.GoogleCloudProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	}
	if len(c.TelegramApiToken) == 0 {
		if len(os.Getenv("TELEGRAM_API_TOKEN")) == 0 {
			log.Fatal("TelegramApiToken config not set!")
		}
		c.TelegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	}
	return nil
}
