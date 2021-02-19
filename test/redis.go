package test

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisPool *redis.Pool
)

func init() {
	NewRedisPool("127.0.0.1:6379")
}

//初始化一个pool
func NewRedisPool(host string) {
	RedisPool = &redis.Pool{
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxIdle: 100,
		//最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
		MaxActive: 100000,
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 120 * time.Second,

		//Dial是创建链接的方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		//TestOnBorrow是一个测试链接可用性的方法
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisQuery(commandName string, args ...interface{}) (reply interface{}, err error) {
	red := RedisPool.Get()
	defer red.Close()

	return red.Do(commandName, args...)
}
