package main

import (
	"github.com/FDDQA/quickrates_bot/internal/bot"
	"github.com/FDDQA/quickrates_bot/internal/config"
)

func main() {
    bot := bot.New()
    bot.Start()
}