package router

import (
	"api/controller"
	"github.com/gin-gonic/gin"
)

//路由基本信息
func Configure(app *gin.Engine) {
	//将整体MVP逻辑进行容器注册
	controllers := &controller.Controllers{}
	controllers.New()

	//通用的路由配置
	v1 := app.Group("/")
	{
		v1.GET("/get/user/:userId", controllers.Hello.GetUserId)
		v1.GET("/get/test/:userId", controllers.Test.GetFirstId)
	}
}
