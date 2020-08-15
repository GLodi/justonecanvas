package rest

import (
	"net/http"

	"github.com/GLodi/justonecanvas/server/internal/api/ws"
	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type canvasHandler struct {
	log *logrus.Logger
	hub *ws.Hub
	uc  canvas.UseCase
}

func NewCanvasHandler(l *logrus.Logger, h *ws.Hub, uc canvas.UseCase) *canvasHandler {
	return &canvasHandler{log: l, hub: h, uc: uc}
}

func (ch *canvasHandler) UpdateRedis() {
	for {
		select {
		case message := <-ch.hub.Store:
			ch.log.Infoln("canvas_handler UpdateRedis:", message)
			// TODO: parse from byte to (pos int, color uint8)
			// ch.uc.Update()
		}

	}
}

func (ch *canvasHandler) Get(ctx *gin.Context) {
	c, err := ch.uc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.log.Errorln("canvas_handler /GET:", err)
		return
	}

	// binary response
	// out := bytes.NewBuffer(nil)
	// if err := gob.NewEncoder(out).Encode(&c); err != nil {
	// 	ch.log.Errorln("canvas_handler /GET: ", err)
	// 	ctx.AbortWithStatusJSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{"err": err.Error()},
	// 	)
	// 	return
	// }

	// ch.log.Println(len(out.Bytes()))

	// ctx.Data(http.StatusOK, "application/x-gob", out.Bytes())

	// JSON response
	ctx.JSON(http.StatusOK, c)

	// ch.log.Infoln("canvas_handler /GET OK")
}

func (ch *canvasHandler) GetWs(ctx *gin.Context) {
	ws.ServeWs(ch.log, ch.hub, ctx.Writer, ctx.Request, ctx.ClientIP())
}

func (ch *canvasHandler) Update(ctx *gin.Context) {
	// would need to make return type of uc.Update in (*Canvas, err)
	// index := ctx.Param("index")
	// val, err := strconv.Atoi(index)
	// if err != nil || val < 0 || val >= 2500 {
	// 	ctx.Status(http.StatusBadRequest)
	// 	ch.log.Errorln("canvas_handler /PUT INDEX RANGE:", err)
	// 	return
	// }

	// color := ctx.Param("color")
	// a, err := strconv.ParseUint(color, 10, 8)
	// if err != nil || a < 0 || a >= 256 {
	// 	ctx.Status(http.StatusBadRequest)
	// 	ch.log.Errorln("canvas_handler /PUT COLOR RANGE:", err)
	// 	return
	// }
	// valc := uint8(a)

	// c, err := ch.uc.Update(val, valc)
	// if len(c.Cells) == 0 || err != nil {
	// 	ctx.Status(http.StatusNoContent)
	// 	ch.log.Errorln("canvas_handler /PUT Update:", err)
	// 	return
	// }

	// ch.log.Infoln("canvas_handler /POST OK")
	// ctx.JSON(http.StatusOK, c)
}
