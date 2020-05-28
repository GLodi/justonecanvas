package canvas

import "github.com/sirupsen/logrus"

type Service interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}

type service struct {
	repo Repository
	l    *logrus.Logger
}

func NewService(r Repository, l *logrus.Logger) Service {
	return &service{
		repo: r,
		l:    l,
	}
}

func (u *service) Get() (*Canvas, error) {
	u.l.Infoln("canvas_service Get()")
	return u.repo.Get()
}

func (u *service) Update(pos int, color uint8) error {
	return u.repo.Update(pos, color)
}
