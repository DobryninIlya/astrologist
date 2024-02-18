package cmd

import (
	"astrologist/internal/app/store/sqlstore"
	"astrologist/internal/bot/commands/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
)

const astrologMessage = `🧙‍♀️ *Натальная карта* - это графическое представление позиций планет на небесной сфере на момент вашего рождения. Она является индивидуальной картой небесного свода и отражает распределение планет по знакам зодиака и домам на момент вашего рождения.✨
*Натальная карта* используется в астрологии для анализа личности, характера, судьбы и потенциальных тенденций жизни человека. Она может помочь понять вашу индивидуальность, сильные и слабые стороны, а также предсказать возможные события и тенденции в вашей жизни на основе астрологических аспектов и влияния планет.💫`

func AstrologCommandProcessor(bot *tgbotapi.BotAPI, store sqlstore.StoreInterface, update tgbotapi.Update) error {
	image, err := os.Open("astrologist_hello.png")
	if err != nil {
		log.Println("Произошла ошибка сохранения результата", err.Error())
	}
	body, err := io.ReadAll(image)
	defer image.Close()
	photo := tgbotapi.FileBytes{
		Name:  "result",
		Bytes: body,
	}
	photoConfig := tgbotapi.NewPhoto(update.Message.Chat.ID, photo)
	photoConfig.Caption = astrologMessage
	//bot.Send(photoConfig)
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, astrologMessage)
	photoConfig.ReplyMarkup = keyboards.NatalInlineKeyboard
	photoConfig.ParseMode = "markdown"
	bot.Send(photoConfig)
	return nil
}
