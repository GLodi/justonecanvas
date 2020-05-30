package canvas

import (
	"github.com/sirupsen/logrus"
)

type UseCase interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}

type usecase struct {
	repo Repository
	l    *logrus.Logger
}

func NewUseCase(r Repository, l *logrus.Logger) UseCase {
	return &usecase{
		repo: r,
		l:    l,
	}
}

func (u *usecase) Get() (*Canvas, error) {
	u.l.Infoln("canvas_usecase Get()")
	return u.repo.Get()
}

func (u *usecase) Update(pos int, color uint8) error {
	u.l.Infoln("canvas_usecase Update()")
	return u.repo.Update(pos, color)
}
