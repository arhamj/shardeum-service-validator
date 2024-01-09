package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := app{}
	app.initConfig()
	app.initValidator()
	app.initDB()
	app.initRepo()
	app.initService()
	app.initScheduler()
	app.initControllers()
	app.initServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.shutdown(ctx); err != nil {
		log.WithError(err).Error("error shutting the app down")
	}
}
