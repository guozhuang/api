package controller

import (
	"github.com/gin-gonic/gin"
)

type TestController struct {
	//
}

func (test *TestController) GetFirstId(c *gin.Context) {
	userId := c.Param("userId")

	userId = BaseModel.TestModel.GetTestInfo(userId)

	c.JSON(200, gin.H{"message": userId})
}
