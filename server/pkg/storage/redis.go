package storage

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewRedis(l *logrus.Logger) *redis.Client {
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	redis := redis.NewClient(&redis.Options{
		Addr:     "redis:" + redis_port,
		Password: redis_password,
		DB:       0,
	})
	pong, err := redis.Ping(context.TODO()).Result()
	if err != nil {
		panic("failed to connect to redis" + err.Error())
	}
	defer redis.Close()
	l.Println(pong, err.Error())
	return redis
}
