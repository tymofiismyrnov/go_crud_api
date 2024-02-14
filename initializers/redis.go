package initializers

import (
	"context"
	"os"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var CTX context.Context

func ConnectoToRedis() {
	CTX = context.Background()
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
