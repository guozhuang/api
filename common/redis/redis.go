package redis

import (
	"api/config"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

//redis操作标准包【进行每个方法的标准化实现】
type Redis struct {
	Pool *redigo.Pool
}

func (redis *Redis) NewPool(serverAddr string) {
	redis.Pool = &redigo.Pool{
		MaxIdle:     config.RedisCommonSetting.MaxIdle, //空闲数
		IdleTimeout: config.RedisCommonSetting.IdleTimeout,
		MaxActive:   config.RedisCommonSetting.MaxActive, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", serverAddr)
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
