package canvas

import "time"

type Canvas struct {
	ID    string     `json:"id"`
	Cells [2500]Cell `json:"cells"`
}

type Cell struct {
	Timestamp time.Time `json:"timestamp"`
	Color     byte      `json:"color"`
}
