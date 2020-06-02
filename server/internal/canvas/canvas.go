package canvas

type Canvas struct {
	ID    string       `json:"id"`
	Cells [2500]uint16 `json:"cells"`
}
