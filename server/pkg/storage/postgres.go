package storage

import (
	"os"

	"github.com/GLodi/justonecanvas/server/pkg/canvas"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func NewPostgres(l *logrus.Logger) (*gorm.DB, error) {
	postgres_port := os.Getenv("POSTGRES_PORT")
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")

	pg, err := gorm.Open("postgres",
		"port="+postgres_port+
			" host="+postgres_host+
			" user="+postgres_user+
			" password="+postgres_password+
			" sslmode=disable")
	if err != nil {
		panic("failed to connect to postgres" + err.Error())
	}

	l.Infoln("Connected to postgres")
	pg.AutoMigrate(&canvas.Canvas{})

	if pg.First(&canvas.Canvas{}).RecordNotFound() {
		c := canvas.Canvas{Cells: [2500]uint8{}}
		pg.Create(&c)
	}
	return pg, nil
}
