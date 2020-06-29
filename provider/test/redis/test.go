package redis

import (
	"api/common/redis"
	"api/config"
)

var testRedis = &redis.Redis{}

func init() {
	testRedis.ServerAddr = config.RedisServerSetting.Test
	testRedis.NewPool()
}

type Provider struct {
	//
}

func (provider *Provider) GetUserName(userId string) string {
	redisValue := testRedis.Get("test")
	return userId + redisValue + " "
}
