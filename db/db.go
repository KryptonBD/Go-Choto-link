package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func CreateClient(dbId int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbId,
	})

	return rdb
}
