package rest

import (
	"net/http"
	"strconv"

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

func (ch *canvasHandler) Get(ctx *gin.Context) {
	c, err := ch.uc.Get()
	if len(c.Cells) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		ch.log.Errorln("canvas_handler /GET:", err)
		return
	}

	// out := bytes.NewBuffer(nil)
	// if err := gob.NewEncoder(out).Encode(&c); err != nil {
	// 	ch.log.Errorln("canvas_handler /GET: ", err)
	// 	ctx.AbortWithStatusJSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{"err": err.Error()},
	// 	)
	// 	return
	// }
	// ctx.Data(http.StatusOK, "application/x-gob", out.Bytes())

	ctx.JSON(http.StatusOK, c)
	ch.log.Infoln("canvas_handler /GET OK")
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// IN ORDER TO TEST FROM BROWSER:
// ws = new WebSocket("ws://localhost:8080/api/v1/canvas/ws");
// ws.onmessage = function (e) {
//     console.log(e.data);
// };

// ws.onopen = function() {
//     ws.send("SomeMessage");
// }

func (ch *canvasHandler) GetWs(ctx *gin.Context) {
	ws.ServeWs(ch.hub, ctx.Writer, ctx.Request)

	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// conn, _ := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	// for {
	// 	msgType, msg, err := conn.ReadMessage()
	// 	if err != nil {
	// 		return
	// 	}

	// 	ch.log.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	// 	msg1 := []byte("prova")

	// 	if err = conn.WriteMessage(msgType, msg1); err != nil {
	// 		return
	// 	}
	// }
}

func (ch *canvasHandler) Update(ctx *gin.Context) {
	index := ctx.Param("index")
	val, err := strconv.Atoi(index)
	if err != nil || val < 0 || val >= 2500 {
		ctx.Status(http.StatusBadRequest)
		ch.log.Errorln("canvas_handler /PUT INDEX RANGE:", err)
		return
	}

	color := ctx.Param("color")
	a, err := strconv.ParseUint(color, 10, 8)
	if err != nil || a < 0 || a >= 256 {
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
