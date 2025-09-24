package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/dachichang/gokit/httpserver"
	"github.com/dachichang/gokit/logger"
	"github.com/sirupsen/logrus"

	"github.com/dachichang/goath-stack-template/internal/interface/api"
	"github.com/dachichang/goath-stack-template/internal/interface/router"
	"github.com/dachichang/goath-stack-template/internal/interface/web"
	_ "github.com/joho/godotenv/autoload"
)

const (
	ShutdownTimeout = 3
)

// @title Goath-Stack
// @version 1.0
// @description Goath-Stack
func main() {
	ctx := context.Background()
	log := logger.NewLogger(os.Getenv("LOG_PATH"), logrus.TraceLevel)

	// create repo ...

	// create service ...

	// api handler
	systemHandler := api.NewSystemHandler()
	// page controller
	webController := web.NewController()

	// route
	routes := router.NewRoutes(
		systemHandler,
		webController,
		log,
	)

	// create server
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := httpserver.NewServer(
		log,
		routes,
		httpserver.WithPort(port),
	)

	// signal catch channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	server.Start()

	<-quit // wait for signal

	ctx, cancel := context.WithTimeout(ctx, time.Duration(ShutdownTimeout)*time.Second)
	defer cancel()

	server.Stop(ctx) // at most block ShutdownTimeout seconds

	log.Info("Graceful shutdown completed")
}
