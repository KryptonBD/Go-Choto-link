package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbId int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbId,
	})

	return rdb
}

func CheckIfShortURLExists(shortUrl string) bool {
	rdb := CreateClient(0)
	defer rdb.Close()

	_, err := rdb.Get(Ctx, shortUrl).Result()

	return err == nil
}
