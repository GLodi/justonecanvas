package rest

import (
	"net/http"

	"github.com/GLodi/justonecanvas/server/pkg/canvas"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type canvasHandler struct {
	l   *logrus.Logger
	svc canvas.Service
}

func NewCanvasHandler(l *logrus.Logger, svc canvas.Service) *canvasHandler {
	return &canvasHandler{l, svc}
}

func (ch *canvasHandler) Get(ctx *gin.Context) {
	ch.l.Infoln("canvas_handler /GET")
	c, err := ch.svc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.l.Errorln("/canvas GET NO CONTENT")
		return
	}
	ch.l.Infoln("canvas_handler /GET OK", c)
	ctx.JSON(http.StatusOK, c)
}
