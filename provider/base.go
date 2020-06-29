package provider

import (
	"api/provider/user"
	"api/utils/inject"
)

//统一对外结构体
type BaseProvider struct {
	User *user.Provider `auto:"userProvider"`
}

func (provider *BaseProvider) New() {
	//动态将相应服务注册在相应结构内，然后挂载到baseProvider内
	inject.Register("baseProvider", provider)
	inject.Inject()
}
