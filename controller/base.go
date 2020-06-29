package controller

import (
	"api/models"
	"api/utils/inject"
)

var BaseModel = &models.BaseModels{}

type Controllers struct {
	Hello *HelloController `auto:"helloController"`
	Test  *TestController  `auto:"testController"`
}

func (ctx *Controllers) New() {
	BaseModel.New()
	inject.Register("baseController", ctx)
	inject.Inject()
}
