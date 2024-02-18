package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	MainKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Натальная карта"),
		),
	)
	NatalInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Рассчитать", "https://t.me/astrologist_robot/natal"),
		),
	)
)
