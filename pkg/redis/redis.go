package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var pool *redis.Pool

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")
	server := fmt.Sprintf("%s:%s", host, port)
	pool = &redis.Pool{
		MaxActive:   100,
		MaxIdle:     100,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// 鉴权
			if pass != "" {
				if _, err := c.Do("AUTH", pass); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}

// 获取实例
func Get() redis.Conn {
	return pool.Get()
}