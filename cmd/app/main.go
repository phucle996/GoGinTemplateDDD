package main

import (
	"authen_service/internal/app"
	"authen_service/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// load config from file .env
	cfg := config.Load()

	// intial application
	application := app.NewApplication(cfg)

	go func() {
		if err := application.Start(cfg); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	application.Stop()
}
