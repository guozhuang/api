package test

import (
	"api/provider/test/redis"
	"api/utils/inject"
)

var provider = &Provider{}

type Provider struct {
	RedisProvider *redis.Provider `auto:"testRedisProvider"`
}

func init() {
	inject.Register("testProvider", provider)
	inject.Inject()
}
