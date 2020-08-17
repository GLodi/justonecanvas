package canvas

import (
	"time"

	"github.com/GLodi/justonecanvas/server/internal/constants"
)

type Canvas struct {
	ID    string                  `json:"id"`
	Cells [constants.Squares]Cell `json:"cells"`
}

type Cell struct {
	Timestamp time.Time `json:"timestamp"`
	Color     byte      `json:"color"`
}
