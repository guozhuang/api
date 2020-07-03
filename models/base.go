package models

import (
	"api/provider"
	"api/utils/inject"
)

var baseProvider *provider.BaseProvider

type BaseModels struct {
	Hello *HelloModel `auto:"helloModel"`
	Test  *TestModel  `auto:"testModel"`
}

func (model *BaseModels) New() {
	baseProvider = provider.GetProvider()
	inject.Register("baseModel", model)
	inject.Inject()
}
