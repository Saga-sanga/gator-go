package main

import (
	"log"
	"os"

	"github.com/saga-sanga/gator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	gatorState := &state{
		cfg: &conf,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: clid <command> [args...]")
	}

	userCommand := command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = cmds.run(gatorState, userCommand)
	if err != nil {
		log.Fatal(err)
	}
}
