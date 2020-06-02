package canvas

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) (*Canvas, error)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.mongo.Database("canvas").Collection("canvas")
	err = collection.FindOne(ctx, bson.M{}).Decode(&c)
	if err != nil {
		r.log.Errorln("canvas_repo Get() FINDONE:", err)
		return nil, err
	}

	return c, nil
}

func (r *repo) Update(pos int, color uint8) (c *Canvas, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.mongo.Database("canvas").Collection("canvas")
	res, err := collection.UpdateMany(ctx,
		bson.M{"id": "1"},
		bson.M{"$set": bson.M{"cells." + strconv.Itoa(pos): color}},
	)
	if err != nil {
		r.log.Errorln("canvas_repo Update() UPDATEONE:", err)
		return nil, err
	}
	r.log.Infoln("canvas_repo Update() UPDATEONE:", res)

	err = collection.FindOne(ctx, bson.M{}).Decode(&c)
	if err != nil {
		r.log.Errorln("canvas_repo Update() FINDONE:", err)
		return nil, err
	}
	return c, nil
}
