package provider

import (
	"api/provider/test"
	"api/provider/user"
	"api/utils/inject"
)

var Base = &BaseProvider{}

//provider标准结构
type BaseProvider struct {
	User *user.Provider `auto:"userProvider"`
	Test *test.Provider `auto:"testProvider"`
}

func (provider *BaseProvider) New() {
	//需要注意挂载provider下的tag需要和内部保持一致
	inject.Register("baseProvider", provider)
	inject.Inject()
}

func GetProvider() *BaseProvider {
	Base.New()
	return Base
}
