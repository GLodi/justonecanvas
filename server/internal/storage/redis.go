package storage

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/GLodi/justonecanvas/server/internal/constants"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRedis(log *logrus.Logger, m *mongo.Client) (client *redis.Client) {
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:" + redis_port,
		Password: redis_password,
		DB:       0,
	})

	c := &canvas.Canvas{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.Database("canvas").Collection("canvas")
	err := collection.FindOne(ctx, bson.D{}).Decode(&c)
	if err != nil {
		log.Errorln("NewRedis CAN'T FIND CANVAS:", err)
	}

	for i := 0; i < constants.Squares; i++ {
		err := client.BitField(ctx, "canvas", "SET", "u4", "#"+strconv.Itoa(i), c.Cells[i].Color).Err()

		if err != nil {
			log.Errorln("NewRedis CAN'T SET FROM CANVAS:", err)
		}
	}

	return client
}
