package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/app"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cryptossApp, err := app.New()
	if err != nil {
		log.Errorf("failed to create new cryptoss server app: %w", err)
	}

	cryptossApp.InitializeRoutes()

	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)

	go func() {
		if err := cryptossApp.Run(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				errChan <- err
			} else {
				close(errChan)
			}
		}
	}()

	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errChan:
		if err != nil {
			log.Errorf("http server was closed with an error: %v", err)
		}
	case <-sigChan:
		log.Info("os signal detected")
	}

	log.Infof("starting graceful shutdown")

	if err := cryptossApp.GracefulShutdown(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
