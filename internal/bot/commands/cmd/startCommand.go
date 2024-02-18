package cmd

import (
	"astrologist/internal/app/store/sqlstore"
	"astrologist/internal/bot/commands/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const startMessage = "Привет! Я бот-астролог. Чтобы рассчитать натальную карту, нажми на кнопку."

func StartCommandProcessor(bot *tgbotapi.BotAPI, store sqlstore.StoreInterface, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)
	msg.ReplyMarkup = keyboards.MainKeyboard
	bot.Send(msg)
	return nil
}
