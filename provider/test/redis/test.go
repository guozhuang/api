package redis

import (
	"api/common/redis"
	"api/config"
	"github.com/gin-gonic/gin"
)

var testRedis redis.IRedis

func init() {
	//判断环境:后续考虑此处的代码标准化，应该是初始化的过程标准化声明到其他位置，并且进行加载统一使用
	if config.AppSetting.RunMode == gin.DebugMode {
		testRedis = &redis.Redis{
			ServerAddr: config.RedisServerSetting.Test,
		}
	} else {
		//todo:
		testRedis = &redis.Cluster{
			ServerAddr: config.RedisServerSetting.Test,
			IsCluster:  true,
		}
	}

	testRedis.NewPool()
}

type Provider struct {
	//
}

func (provider *Provider) GetUserName(userId string) string {
	testRedis.Set("test", "new set")
	redisValue := testRedis.Get("test")
	return userId + redisValue + " "
}
