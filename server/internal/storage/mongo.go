package storage

import (
	"context"
	"os"
	"time"

	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/GLodi/justonecanvas/server/internal/constants"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo(log *logrus.Logger) (client *mongo.Client, err error) {
	mongo_username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongo_password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	opt := options.Client().ApplyURI(
		"mongodb://" + mongo_username +
			":" + mongo_password +
			"@mongo:27017")
	client, err = mongo.NewClient(opt)
	if err != nil {
		log.Errorln("NewMongo: CAN'T CONNECT", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Errorln("NewMongo: COULDN'T CONNECT", err)
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Errorln("NewMongo: NO PING", err)
		return nil, err
	}

	collection := client.Database("canvas").Collection("canvas")
	err = collection.FindOne(ctx, bson.D{}).Err()
	if err != nil {
		log.Errorln("NewMongo CAN'T FIND CANVAS:", err)

		c := &canvas.Canvas{ID: "1"}
		cell := canvas.Cell{
			Color: 0,
		}
		for i := 0; i < constants.Squares; i++ {
			c.Cells[i] = cell
		}

		collection.InsertOne(ctx, c)
	}

	return client, nil
}
