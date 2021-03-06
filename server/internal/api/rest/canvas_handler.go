package rest

import (
	"net/http"
	"time"

	"github.com/GLodi/justonecanvas/server/internal/api/ws"
	"github.com/GLodi/justonecanvas/server/internal/canvas"
	"github.com/GLodi/justonecanvas/server/internal/constants"
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
			ch.uc.Update(int(message[1])*constants.SquarePerRow+int(message[2]), message[0])
		}

	}
}

func (ch *canvasHandler) UpdateMongo() {
	for range time.Tick(5 * time.Second) {
		ch.uc.LiveUpdate()
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
