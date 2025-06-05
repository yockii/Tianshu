package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/yockii/Tianshu/pkg/config"
)

var Pool *redis.Pool

func InitRedis() {
	cfg := config.Cfg.Redis
	Pool = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
				redis.DialPassword(cfg.Password),
				redis.DialDatabase(cfg.DB),
			)
		},
	}
}
