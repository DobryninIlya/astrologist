package commands

import (
	"astrologist/internal/app/store/sqlstore"
	"astrologist/internal/bot/commands/cmd"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//cmd "astrologist/internal/bot/commands/cmd"
)

var (
	CommandList Commands
)

func init() {
	CommandList = Commands{Commands: make([]Command, 0)}
	RegisterCommand([]string{"/start", "start"}, cmd.StartCommandProcessor, "Start command")
	RegisterCommand([]string{"натальная карта", "рассчитать", "расклад"}, cmd.AstrologCommandProcessor, "Start command")

}

type Commands struct {
	Commands []Command
}

type Command struct {
	Keys        []string
	Processor   func(bot *tgbotapi.BotAPI, store sqlstore.StoreInterface, update tgbotapi.Update) error
	Description string
}

func RegisterCommand(keys []string, processor func(bot *tgbotapi.BotAPI, store sqlstore.StoreInterface, update tgbotapi.Update) error, description string) {
	CommandList.Commands = append(CommandList.Commands, Command{
		Keys:        keys,
		Processor:   processor,
		Description: description,
	})
}
