package rest

import (
	"net/http"

	"github.com/GLodi/justonecanvas/server/pkg/canvas"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type canvasHandler struct {
	l  *logrus.Logger
	uc canvas.UseCase
}

func NewCanvasHandler(l *logrus.Logger, uc canvas.UseCase) *canvasHandler {
	return &canvasHandler{l, uc}
}

func (ch *canvasHandler) Get(ctx *gin.Context) {
	ch.l.Infoln("canvas_handler /GET")
	c, err := ch.uc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.l.Errorln("canvas_handler /GET NO CONTENT")
		return
	}
	ch.l.Infoln("canvas_handler /GET OK", c)
	ctx.JSON(http.StatusOK, c)
}
