package main

import (
	"log"

	"github.com/hagit4/goph-keeper/internal/keeper/server"
)

func main() {
	s, err := server.NewKeeperServer()
	if err != nil {
		log.Fatalln(err)
	}
	if err = s.Serve(); err != nil {
		log.Fatalln(err)
	}
}
