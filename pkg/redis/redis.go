package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/xiam/to"
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
		MaxActive:   to.Int(os.Getenv("REDIS_MAX_CONNECTIONS")),
		MaxIdle:     to.Int(os.Getenv("REDIS_MAX_IDLE_CONNECTIONS")),
		IdleTimeout: time.Duration(to.Int(os.Getenv("REDIS_MAX_IDLE_TIMEOUT"))) * time.Second,
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

// GetInstance 获取 Redis 实例
func GetInstance() redis.Conn {
	return pool.Get()
}
