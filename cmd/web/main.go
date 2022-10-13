package main

import (
	"backend-service/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs.Initialize()
	r := gin.Default()
	r.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "version": viper.GetString("version")})
	})
	r.Run(":8100")
}
