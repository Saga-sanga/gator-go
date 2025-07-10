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
	fmt.Printf("Read config: %+v\n", conf)

	err = conf.SetUser("sanga")
	if err != nil {
		log.Fatal(err)
	}

	conf, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v", conf)
}
