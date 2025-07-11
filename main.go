package main

import (
	"log"

	"github.com/saga-sanga/gator-go/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	gatorState := state{
		config: &conf,
	}

	startGator(&gatorState)
}
