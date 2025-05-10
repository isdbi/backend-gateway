package redis

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

func Connect() *redis.Client {
	rdb := *redis.NewClient(
		&redis.Options{
      Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Failed to connect to redis",err.Error())
		return nil
	}

	return &rdb
}
