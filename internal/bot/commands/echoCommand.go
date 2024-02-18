package commands

import (
	"astrologist/internal/app/store/sqlstore"
	"astrologist/internal/bot/commands/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	RegisterCommand([]string{"/start", "start"}, startCommandProcessor, "Start command")
}

func startCommandProcessor(bot *tgbotapi.BotAPI, store sqlstore.StoreInterface, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = keyboards.MainKeyboard
	bot.Send(msg)
	return nil
}
