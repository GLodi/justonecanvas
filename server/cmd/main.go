package main

import (
	"fmt"
	"os"

	"github.com/GLodi/justonecanvas/server/internal/api"
	"github.com/GLodi/justonecanvas/server/internal/di"
	"github.com/gin-gonic/gin"
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

	var l *logrus.Logger
	d.Invoke(func(log *logrus.Logger) {
		l = log
	})
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	svr := api.NewServer(g, d, l)
	svr.MapRoutes()

	return svr.Start()
}
