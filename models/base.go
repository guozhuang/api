package models

import "api/utils/inject"

//model层同样不进行分层：实践中很多标准库也是这种结构
type BaseModels struct {
	HelloModel *HelloModel  `auto:"helloModel"`
	TestModel  *TestModel   `auto:"testModel"`
}

func (m *BaseModels) New () {
	inject.Register("baseModel", m)
	inject.Inject()
}
