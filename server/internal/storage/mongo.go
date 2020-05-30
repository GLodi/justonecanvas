package storage

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() (*mongo.Client, error) {
	mongo_username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongo_pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	opt := options.Client().ApplyURI(
		"mongodb://" + mongo_username +
			":" + mongo_pass +
			"@localhost:27017")
	client, err := mongo.Connect(ctx, opt)
	defer cancel()
	return client, err
}
