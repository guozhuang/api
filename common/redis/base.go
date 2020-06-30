package redis

//统一操作类似redis的方法集
type IRedis interface {
	NewPool()
	Get(string) string
	Set(string, string)
	MGet([]string) map[string]string
}
