package bot

import (
	"astrologist/internal/bot/commands"
	_ "astrologist/internal/bot/commands"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"strings"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	log     *logrus.Logger
	ctx     context.Context
}

func NewBot(ctx context.Context, log *logrus.Logger, token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return &Bot{
		bot:     bot,
		updates: updates,
		ctx:     ctx,
		log:     log,
	}, nil
}

func (b *Bot) HandleUpdates() {
	for update := range b.updates {
		if update.Message == nil {
			continue
		}
		b.log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message == nil {
			continue // Обрабатываем только сообщения с текстом
		}
		lowerText := strings.ToLower(update.Message.Text)
		for _, cmd := range commands.CommandList.Commands {
			for _, key := range cmd.Keys {
				if key == lowerText {
					err := cmd.Processor(b.bot, nil, update)
					if err != nil {
						b.log.Error(err)
					}
				}
			}
		}
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//b.bot.Send(msg)
	}
}
