package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/FDDQA/quickrates_bot/internal/bot"
)

func (b *Bot) handleStart(msg *tgbotapi.Message) {
    text := `
    👋 Welcome to *QuickRates Bot*!
    Commands:
    /start - Show this message
    /add [pair] - Track a currency pair (e.g., /add BTC-USDT)
    `
    reply := tgbotapi.NewMessage(msg.Chat.ID, text)
    reply.ParseMode = "Markdown"
    b.api.Send(reply)
}