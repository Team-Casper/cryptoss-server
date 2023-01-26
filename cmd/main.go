package main

import (
	"github.com/team-casper/cryptoss-server/server"
	"os"
)

func main() {
	srv := server.New()

	if err := srv.Run(); err != nil {
		os.Exit(1)
	}
}
