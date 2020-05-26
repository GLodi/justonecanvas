package canvas

import "github.com/jinzhu/gorm"

type Canvas struct {
	gorm.Model
	Cells [2500]uint8 `json:"cells" gorm:"column:cells"`
}

type Repository interface {
	Get() (*Canvas, error)
	Update(pos int, color uint8) error
}
