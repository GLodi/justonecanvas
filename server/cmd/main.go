package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GLodi/justonecanvas/server/internal/api"
	"github.com/GLodi/justonecanvas/server/internal/api/rest"
	"github.com/GLodi/justonecanvas/server/internal/di"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	pprof.Register(g)

	d := di.BuildContainer()

	d.Invoke(func(redClient *redis.Client) {
		g.Use(rest.NewRateLimiterMiddleware(redClient, "general", 30, 60*time.Second))
	})

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
