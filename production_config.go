//+build !test

package main

import (
	"github.com/gin-gonic/gin"
	"telegram_bot_api/config"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.InitConfig()
}