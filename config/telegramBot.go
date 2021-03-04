package config

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

type configTelegramBot struct {
	BotToken string
	BaseURL string
}

var ConfigTelegramBot configTelegramBot

func initTelegramBotConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		ConfigTelegramBot = configTelegramBot{"", ""}
	case gin.DebugMode:
		ConfigTelegramBot = configTelegramBot{os.Getenv("TELEGRAM_API_TOKEN"), os.Getenv("BaseURL")}
	case gin.TestMode:
		ConfigTelegramBot = configTelegramBot{"", ""}
	}
}

func TelegramBotInit() (bot *tgbotapi.BotAPI, err error){
	bot, _ = tgbotapi.NewBotAPI(ConfigTelegramBot.BotToken)
	url := ConfigTelegramBot.BaseURL + ConfigTelegramBot.BotToken
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	return bot, err
}