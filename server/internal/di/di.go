package di

import (
	"github.com/GLodi/justonecanvas/server/internal/api/ws"
	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/GLodi/justonecanvas/server/internal/storage"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// logger
	container.Provide(logrus.New)
	container.Provide(ws.NewHub)

	// storage
	container.Provide(storage.NewMongo)
	container.Provide(storage.NewRedis)

	// canvas
	container.Provide(canvas.NewRepo)
	container.Provide(canvas.NewUseCase)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
