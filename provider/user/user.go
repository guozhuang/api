package user

import (
	"api/provider/user/mongo"
	"api/provider/user/redis"
	"api/utils/inject"
)

var provider = &Provider{}

type Provider struct {
	RedisProvider *redis.Provider `auto:"userRedisProvider"`
	MongoProvider *mongo.Provider `auto:"userMongoProvider"`
}

func init() {
	inject.Register("userProvider", provider)
	inject.Inject()
}
