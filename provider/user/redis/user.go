package redis

func init() {
	//进行配置化获取，并且进行连接池的初始化，用于整体使用时，获取连接
}

type Provider struct {
	//
}

func (provider *Provider) GetUserName(userId string) string {
	return userId + "redis "
}
