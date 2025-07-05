package handlers

import (
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/FDDQA/quickrates_bot/internal/bot"

)

func (b *Bot) handleAdd(msg *tgbotapi.Message) {
    args := strings.Split(msg.CommandArguments(), "-")
    if len(args) != 2 {
        b.api.Send(tgbotapi.NewMessage(
            msg.Chat.ID,
            "❌ Invalid format. Use: /add BTC-USDT",
        ))
        return
    }

    // TODO: Сохранить пару в БД (пока просто отвечаем)
    pair := strings.ToUpper(args[0]) + "-" + strings.ToUpper(args[1])
    reply := tgbotapi.NewMessage(
        msg.Chat.ID,
        "✅ Added pair: " + pair,
    )
    b.api.Send(reply)
}