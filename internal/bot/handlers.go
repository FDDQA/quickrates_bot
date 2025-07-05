package bot

import (
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStart(msg *tgbotapi.Message) {
	text := "👋 Добро пожаловать в QuickRates!\n\n" +
		"Доступные команды:\n" +
		"/start - показать это сообщение\n" +
		"/add - добавить валютную пару\n" +
		"/remove - удалить пару\n" +
		"/list - показать все пары"

	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	b.api.Send(reply)
}

func (b *Bot) handleAdd(msg *tgbotapi.Message) {
	args := strings.Split(strings.ToUpper(msg.CommandArguments()), "-")
	if len(args) != 2 {
		reply := tgbotapi.NewMessage(msg.Chat.ID, "❌ Неверный формат. Используйте: /add BTC-RUB")
		b.api.Send(reply)
		return
	}

	pair := args[0] + "-" + args[1]
	reply := tgbotapi.NewMessage(msg.Chat.ID, "✅ Пара добавлена: "+pair)
	b.api.Send(reply)
	
}