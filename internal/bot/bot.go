package bot

import (
	"log"

	"github.com/FDDQA/quickrates_bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func New() *Bot {
	cfg := config.Load()
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized as %s", api.Self.UserName)
	return &Bot{api: api}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start":
			b.handleStart(update.Message)
		case "add":
			b.handleAdd(update.Message)
		default:
			b.api.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"🚀 Unknown command. Use /start",
			))
		}
	}
}