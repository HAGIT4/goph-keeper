package main

import (
	"log"

	"github.com/hagit4/goph-keeper/internal/agent/client"
)

func main() {
	c, err := client.NewAgentClient()
	if err != nil {
		log.Fatal(err)
	}
	c.Execute()
}
