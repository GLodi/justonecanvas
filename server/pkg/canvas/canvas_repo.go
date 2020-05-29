package canvas

import (
	"context"

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
	l     *logrus.Logger
}

func NewRepo(mongo *mongo.Client, redis *redis.Client, l *logrus.Logger) Repository {
	return &repo{
		mongo: mongo,
		redis: redis,
		l:     l,
	}
}

func (r *repo) Get() (c *Canvas, err error) {
	// once you have both redis and pg, check redis first
	r.l.Infoln("canvas_repo Get()")

	c = &Canvas{}

	collection := r.mongo.Database("test").Collection("trainers")
	err = collection.FindOne(context.TODO(), bson.D{{}}).Decode(&c)
	if err != nil {
		r.l.Errorln("canvas_repo Get() didn't find")
		return nil, err
	}

	return c, nil
}

func (r *repo) Update(pos int, color uint8) error {
	r.l.Infoln("canvas_repo Update()")
	return nil
}
