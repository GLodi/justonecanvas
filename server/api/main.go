package main

import (
	"fmt"
	"os"

	"github.com/GLodi/justonecanvas/server/api/server"
	"github.com/GLodi/justonecanvas/server/pkg/di"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
func run() error {
	g := gin.Default()
	d := di.BuildContainer()
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	svr := server.NewServer(g, d, l)
	svr.MapRoutes()

	return svr.Start()
}
