package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"telegram_bot_api/config"
	"telegram_bot_api/util/log"
)

func ＷebHookHandler(c *gin.Context) {
	defer c.Request.Body.Close()
	var update tgbotapi.Update
	if bytes, err := ioutil.ReadAll(c.Request.Body); err != nil {
		log.Error(err)
		return
	} else if err = json.Unmarshal(bytes, &update); err != nil {
		log.Error(err)
		return
	}
	SendMessage(update, update.Message.Text)
}


func SendMessage(update tgbotapi.Update, message string){
	bot, _ := config.TelegramBotInit()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}


func SendImage(update tgbotapi.Update, fileUrl string){
	bot, _ := config.TelegramBotInit()
	msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, fileUrl)
	bot.Send(msg)
}
