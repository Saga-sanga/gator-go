package main

import (
	"fmt"
	"os"

	"github.com/saga-sanga/gator-go/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("The login handler expects a single argument, the username")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set", cmd.arguments[0])

	return nil
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	currentCmd, ok := c.handlers[cmd.name]
	if ok {
		err := currentCmd(s, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func startGator(s *state) {
	gatorCommands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	gatorCommands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Need more arguments")
		os.Exit(1)
	}

	userCommand := command{
		name:      args[1],
		arguments: args[2:],
	}

	err := gatorCommands.run(s, userCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
