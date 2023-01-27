package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/team-casper/cryptoss-server/server"
	"os"
)

func main() {
	srv := server.New()

	if err := srv.Run(); err != nil {
		log.Errorf("server is closing: %v", err)
		os.Exit(1)
	}
}
