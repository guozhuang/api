package cmd

import (
	"api/config"
	"api/router"
	"github.com/gin-gonic/gin"
)

//此处进行gin依赖的标准容器
func Init() *gin.Engine {
	//可否将配置信息也同步到相应的结构体中，使用相应的方法才能访问
	gin.SetMode(config.AppSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger()) //gin的中间层的使用
	r.Use(gin.Recovery())

	//分离的各层通过container来进行注册，然后再路由初始化，最后返回gin

	router.Configure(r)

	return r
}
