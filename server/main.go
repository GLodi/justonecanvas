package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres",
		"port="+os.Getenv("POSTGRES_PORT")+
			" host="+os.Getenv("POSTGRES_DB")+
			" user="+os.Getenv("POSTGRES_USER")+
			" password="+os.Getenv("POSTGRES_PASSWORD")+
			" sslmode=disable")

	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	fmt.Println("prova")

	defer db.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
