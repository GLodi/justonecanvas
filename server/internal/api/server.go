package api

import (
	"github.com/GLodi/justonecanvas/server/internal/api/rest"
	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
	logger *logrus.Logger
}

func NewServer(e *gin.Engine, c *dig.Container, l *logrus.Logger) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (ds *dserver) Start() error {
	return ds.router.Run()
}

func (ds *dserver) MapRoutes() {
	apiV1 := ds.router.Group("api/v1")
	ds.canvasRoutes(apiV1)
}

func (ds *dserver) canvasRoutes(api *gin.RouterGroup) {
	canvasRoutes := api.Group("/canvas")
	{
		var cs canvas.UseCase
		ds.cont.Invoke(func(s canvas.UseCase) {
			cs = s
		})

		ch := rest.NewCanvasHandler(ds.logger, cs)

		canvasRoutes.GET("/", ch.Get)
		canvasRoutes.GET("/ws", ch.GetWs)
		canvasRoutes.PUT("/:index/color/:color", ch.Update)
	}
}
