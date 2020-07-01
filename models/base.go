package models

import (
	"api/provider"
	"api/utils/inject"
)

//editor同样需要使用provider//形成一个标准的获取provider方法
var baseProvider *provider.BaseProvider

type BaseModels struct {
	Hello *HelloModel `auto:"helloModel"`
	Test  *TestModel  `auto:"testModel"`
}

func (m *BaseModels) New() {
	baseProvider = provider.GetProvider()
	inject.Register("baseModel", m)
	inject.Inject()
}
