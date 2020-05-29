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
	container.Provide(storage.NewRedis)
	container.Provide(storage.NewMongo)

	// canvas
	container.Provide(canvas.NewRepo)
	container.Provide(canvas.NewUseCase)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
