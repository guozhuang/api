package mongo

//mongo基础加载
type Provider struct {
	//
}

func (provider *Provider) GetMongoName(userId string) string {
	//进行mongo包处理
	return userId + "mongo"
}
