package canvas

import (
	e "github.com/GLodi/justonecanvas/server/pkg/errors"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) Get() (c *Canvas, err error) {
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
