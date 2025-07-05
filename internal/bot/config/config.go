package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    TelegramToken string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    return &Config{
        TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
    }
}
