package storage

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(l *logrus.Logger) (*mongo.Client, error) {
	mongo_username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongo_pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opt := options.Client().ApplyURI(
		"mongodb://" + mongo_username +
			":" + mongo_pass +
			"@localhost:27017")
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic("failed to connect to mongo" + err.Error())
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		l.Errorln(err)
	}

	l.Infoln("Connected to mongo")

	return client, nil
}
