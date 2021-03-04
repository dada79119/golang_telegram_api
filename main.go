package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"telegram_bot_api/config"
	"telegram_bot_api/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err == nil {
		gin.SetMode(gin.DebugMode)
		config.InitConfig()
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST","GET"},
		AllowHeaders:    []string{"Content-Type", "token"},

	}))
	router.POST("/telegram_bot", handler.ＷebHookHandler)
	fmt.Println(config.ServerConfig.Port)
	_ = router.Run(":" + config.ServerConfig.Port)
}