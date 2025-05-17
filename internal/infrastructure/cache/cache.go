package cache

import (
	"github.com/redis/go-redis/v9"
)

func NewRedis(addr string) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	}), nil
}