package canvas

import (
	e "github.com/GLodi/justonecanvas/server/pkg/errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}

type repo struct {
	DB *gorm.DB
	l  *logrus.Logger
}

func NewRepo(db *gorm.DB, l *logrus.Logger) Repository {
	return &repo{
		DB: db,
		l:  l,
	}
}

func (r *repo) Get() (c *Canvas, err error) {
	// once you have both redis and pg, check redis first
	r.l.Infoln("canvas_repo Get()")
	c = &Canvas{}
	result := r.DB.First(c)

	switch result.Error {
	case nil:
		return c, nil
	case gorm.ErrRecordNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, e.ErrDatabase
	}
}

func (r *repo) Update(pos int, color uint8) error {
	c := &Canvas{}
	res1 := r.DB.First(c)

	switch res1.Error {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return e.ErrNotFound
	}

	c.Cells[pos] = color
	res2 := r.DB.Save(c)

	switch res2.Error {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return e.ErrNotFound
	default:
		return e.ErrDatabase
	}
}
