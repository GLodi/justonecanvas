package canvas

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}

type repo struct {
	mongo *mongo.Client
	redis *redis.Client
	log   *logrus.Logger
}

func NewRepo(mongo *mongo.Client, redis *redis.Client, l *logrus.Logger) Repository {
	return &repo{
		mongo: mongo,
		redis: redis,
		log:   l,
	}
}

func (r *repo) Get() (c *Canvas, err error) {
	// once you have both redis and pg, check redis first
	r.log.Infoln("canvas_repo Get()")

	c = &Canvas{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.mongo.Database("canvas").Collection("canvas")
	err = collection.FindOne(ctx, bson.D{{}}).Decode(&c)
	if err != nil {
		r.log.Errorln("canvas_repo Get() FINDONE ERROR:", err)
		return nil, err
	}

	return c, nil
}

func (r *repo) Update(pos int, color uint8) error {
	r.log.Infoln("canvas_repo Update()")
	return nil
}
