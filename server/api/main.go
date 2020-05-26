package main

import (
	"os"

	"github.com/GLodi/justonecanvas/server/pkg/canvas"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func dbInit(port, host, user, password string) (*gorm.DB, error) {
	return gorm.Open("postgres",
		"port="+port+
			" host="+host+
			" user="+user+
			" password="+password+
			" sslmode=disable")
}

func redisInit(port, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:" + port,
		Password: password,
		DB:       0,
	})
}

func main() {
	postgres_port := os.Getenv("POSTGRES_PORT")
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")

	pg, err := dbInit(postgres_port, postgres_host, postgres_user, postgres_password)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	defer pg.Close()
	log.Println("Connected to the database")

	pg.AutoMigrate(&canvas.Canvas{})

	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	redis := redisInit(redis_port, redis_password)
	pong, err := redis.Ping().Result()
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	log.Println(pong, err)
}
