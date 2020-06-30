package redis

//对redis的集群操作【需要对节点上的可用性检查和自动更新】
import (
	redis "github.com/chasex/redis-go-cluster"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type Cluster struct {
	Pool       *redigo.Pool
	ServerAddr string
	IsCluster  bool //标记是否是集群
}

func (cluster *Cluster) NewPool() {
	redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
}

func (cluster *Cluster) Exec() {
	//
}

func (cluster *Cluster) Get(key string) string {
	return key
}

func (cluster *Cluster) Set(key, value string) {

}

func (cluster *Cluster) MGet(keys []string) map[string]string {
	result := make(map[string]string)

	return result
}
