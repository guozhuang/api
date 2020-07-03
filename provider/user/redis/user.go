package redis

import (
	"api/common/redis"
	"api/config"
)

var userRedis = &redis.Redis{}

func init() {
	userRedis.ServerAddr = config.RedisServerSetting.User
	userRedis.NewPool()
}

type Provider struct {
	//
}

func (provider *Provider) GetUserName(userId string) string {
	redisValue := userRedis.Get("test")

	keys := make([]string, 0)
	keys = append(keys, "test")
	keys = append(keys, "hello")
	data := userRedis.MGet(keys)
	for key, value := range data {
		userId += "key=" + key + " ,value=" + value
	}
	return userId + redisValue + " "
}

func (provider *Provider) GetUserInfo(userId string) string {
	redisValue := userRedis.Get(userId)

	return redisValue
}

func (provider *Provider) SetUserInfo(userId, data string) {
	userRedis.Set(userId, data)
}
