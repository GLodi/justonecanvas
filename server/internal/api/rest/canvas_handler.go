package rest

import (
	"net/http"

	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type canvasHandler struct {
	log *logrus.Logger
	uc  canvas.UseCase
}

func NewCanvasHandler(l *logrus.Logger, uc canvas.UseCase) *canvasHandler {
	return &canvasHandler{l, uc}
}

func (ch *canvasHandler) Get(ctx *gin.Context) {
	ch.log.Infoln("canvas_handler /GET")
	c, err := ch.uc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.log.Errorln("canvas_handler /GET:", err)
		return
	}
	ch.log.Infoln("canvas_handler /GET OK")
	ctx.JSON(http.StatusOK, c)
}
