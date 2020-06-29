package mongo

//mongo基础加载
type Provider struct {
	//
}

func (provider *Provider) GetMongoName(userId string) string {
	return userId + "mongo"
}
