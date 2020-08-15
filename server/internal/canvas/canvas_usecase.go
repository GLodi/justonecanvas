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
	log  *logrus.Logger
}

func NewUseCase(r Repository, l *logrus.Logger) UseCase {
	return &usecase{
		repo: r,
		log:  l,
	}
}

func (u *usecase) Get() (*Canvas, error) {
	return u.repo.Get()
}

func (u *usecase) Update(pos int, color uint8) error {
	return u.repo.Update(pos, color)
}
