package cmd

import (
	"api/config"
	"api/router"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(config.AppSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger()) //gin的中间层的使用
	r.Use(gin.Recovery())

	router.Configure(r)

	return r
}
