package models

type HelloModel struct {
	//
}

func (hello *HelloModel) GetHelloData(id string) string {
	id += "aaaa"

	//model需要各自层面实现对应provider的实例化，在当前包内有效还是全局有效，都需要进行判断
	//以及连接池的初始化
	id = baseProvider.User.RedisProvider.GetUserName(id)

	id = baseProvider.Test.RedisProvider.GetUserName(id)

	id = baseProvider.User.MongoProvider.GetMongoName(id)

	return id
}
