package redis

import (
	"api/config"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

//redis操作标准包【进行每个方法的标准化实现】
type Redis struct {
	Pool       *redigo.Pool
	ServerAddr string
}

func (redis *Redis) NewPool() {
	//todo:需要考虑服务的ip切换问题，所以应该在provider内封装一个对应更新的方法【应该完善到redis结构体中：关联相应的provider】
	redis.Pool = &redigo.Pool{
		MaxIdle:     config.RedisCommonSetting.MaxIdle, //空闲数
		IdleTimeout: config.RedisCommonSetting.IdleTimeout,
		MaxActive:   config.RedisCommonSetting.MaxActive, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", redis.ServerAddr)
			//因为如果该服务进行切换的话，仅仅是初始化的时候拿到连接池，是不足以保证后续服务的稳定性
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//统一代理处理
func (redis *Redis) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redis.Pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	result, err := con.Do(cmd, parmas...)

	return result, err
}

//从这里开始进行redis操作基础方法扩展
func (redis *Redis) Get(key string) string {
	data, err := redigo.String(redis.Exec("get", key))
	if err != nil {
		fmt.Println(err) //todo:后续将这里的输出标准化到中间件中
		return ""
	}

	return data
}

func (redis *Redis) Set(key, value string) {
	_, err := redigo.String(redis.Exec("set", key, value))
	if err != nil {
		fmt.Println(err)
	}
}

//不定长mget
func (redis *Redis) MGet(keys []string) map[string]string {
	result := make(map[string]string)

	args := make([]interface{}, 0)
	for i := 1; i < len(keys); i++ {
		args = append(args, keys[i])
	}

	reply, _ := redigo.Strings(redis.Exec("MGET", keys[0], args...))

	for i := 0; i < len(reply); i++ {
		result[keys[i]] = reply[i]
	}

	return result
}

func (redis *Redis) MSet(data map[string]string) {
	//
}
