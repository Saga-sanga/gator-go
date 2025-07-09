package main

import (
	"fmt"
	"log"

	"github.com/saga-sanga/gator-go/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = conf.SetUser("sanga")
	if err != nil {
		log.Fatal(err)
	}

	conf, err = config.Read()
	fmt.Printf("Config: %v", conf)
}
