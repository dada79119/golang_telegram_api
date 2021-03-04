//+build test

package main

import (
	"api/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
	config.InitConfig()
}