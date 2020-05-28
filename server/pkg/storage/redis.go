package storage

import (
	"os"

	"github.com/go-redis/redis"
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
	pong, err := redis.Ping().Result()
	if err != nil {
		panic("failed to connect to redis" + err.Error())
	}
	defer redis.Close()
	l.Println(pong, err.Error())
	return redis
}
