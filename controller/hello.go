package controller

import (
	"github.com/gin-gonic/gin"
)

//控制器不需要分层，将对应的model初始化在当前包内
type HelloController struct {
	//
}

func (hello *HelloController) GetUserId(c *gin.Context) {
	userId := c.Param("userId")

	userId = BaseModel.Hello.GetHelloData(userId)

	c.JSON(200, gin.H{"message": userId})
}
