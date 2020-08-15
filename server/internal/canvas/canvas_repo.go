package canvas

import (
	"context"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
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

var args []interface{}

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

	// collection := r.mongo.Database("canvas").Collection("canvas")
	// err = collection.FindOne(ctx, bson.M{}).Decode(&c)
	// if err != nil {
	// 	r.log.Errorln("canvas_repo Get() FINDONE:", err)
	// 	return nil, err
	// }

	c = &Canvas{}

	if len(args) == 0 {
		r.log.Infoln("ECCOMIT")
		for i := 0; i < 2500; i++ {
			args = append(args, "GET")
			args = append(args, "u4")
			args = append(args, "#"+strconv.Itoa(i))
		}
	}

	val, err := r.redis.BitField(ctx, "canvas", args...).Result()
	if err != nil {
		r.log.Errorln("canvas_repo Get() GETRANGE:", err)
		return nil, err
	}

	b := make([]byte, 8)
	for i := 0; i < 2500; i++ {
		binary.LittleEndian.PutUint64(b, uint64(val[i]))
		c.Cells[i].Color = b[0]

	}

	return c, nil
}

func (r *repo) Update(pos int, color uint8) (err error) {
	// would need to make return type of uc.Update in (*Canvas, err)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// collection := r.mongo.Database("canvas").Collection("canvas")
	// _, err = collection.UpdateMany(ctx,
	// 	bson.M{"id": "1"},
	// 	bson.M{"$set": bson.M{"cells." + strconv.Itoa(pos): color}},
	// )
	// if err != nil {
	// 	r.log.Errorln("canvas_repo Update() UPDATEONE:", err)
	// 	return err
	// }

	// err = collection.FindOne(ctx, bson.M{}).Decode(&c)
	// if err != nil {
	// 	r.log.Errorln("canvas_repo Update() FINDONE:", err)
	// 	return err
	// }
	// return nil

	return nil
}
