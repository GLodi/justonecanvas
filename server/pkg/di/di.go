package di

import (
	"github.com/GLodi/justonecanvas/server/pkg/canvas"
	"github.com/GLodi/justonecanvas/server/pkg/storage"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// logger
	container.Provide(logrus.New)

	// storage
	container.Provide(storage.NewPostgres)
	container.Provide(storage.NewRedis)

	// canvas
	container.Provide(canvas.NewRepo)
	container.Provide(canvas.NewService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
