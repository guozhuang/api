package models

import (
	"api/provider"
	"api/utils/inject"
)

var baseProvider = &provider.BaseProvider{}

//model层同样不进行分层：实践中很多标准库也是这种结构
type BaseModels struct {
	Hello *HelloModel `auto:"helloModel"`
	Test  *TestModel  `auto:"testModel"`
}

func (m *BaseModels) New() {
	baseProvider.New()
	inject.Register("baseModel", m)
	inject.Inject()
}
