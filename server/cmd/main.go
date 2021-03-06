package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GLodi/justonecanvas/server/internal/api"
	"github.com/GLodi/justonecanvas/server/internal/api/rest"
	"github.com/GLodi/justonecanvas/server/internal/di"
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
	// g.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour}))
	// pprof.Register(g)

	d := di.BuildContainer()

	// HACK: rest rate limiter. comment for artillery
	d.Invoke(func(redClient *redis.Client) {
		g.Use(rest.NewRateLimiterMiddleware(redClient, "general", 100, 60*time.Second))
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
