package canvas

import (
	"github.com/GLodi/justonecanvas/server/internal/constants"
)

type Canvas struct {
	ID    string                  `json:"id"`
	Cells [constants.Squares]Cell `json:"cells"`
}

type Cell struct {
	Color byte `json:"color"`
}
