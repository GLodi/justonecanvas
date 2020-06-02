package rest

import (
	"net/http"
	"strconv"

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
	c, err := ch.uc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.log.Errorln("canvas_handler /GET:", err)
		return
	}

	ch.log.Infoln("canvas_handler /GET OK")
	ctx.JSON(http.StatusOK, c)
}

func (ch *canvasHandler) Update(ctx *gin.Context) {
	index := ctx.Param("index")
	val, err := strconv.Atoi(index)
	if err != nil || val < 0 || val > 2499 {
		ctx.Status(http.StatusBadRequest)
		ch.log.Errorln("canvas_handler /PUT INDEX RANGE:", err)
		return
	}

	color := ctx.Param("color")
	a, err := strconv.ParseUint(color, 10, 8)
	if err != nil || a < 0 || a > 255 {
		ctx.Status(http.StatusBadRequest)
		ch.log.Errorln("canvas_handler /PUT COLOR RANGE:", err)
		return
	}
	valc := uint8(a)

	c, err := ch.uc.Update(val, valc)
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.log.Errorln("canvas_handler /PUT Update:", err)
		return
	}

	ch.log.Infoln("canvas_handler /POST OK")
	ctx.JSON(http.StatusOK, c)
}
