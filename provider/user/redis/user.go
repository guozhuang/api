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
	return userId + redisValue + " "
}
