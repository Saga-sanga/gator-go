package main

import (
	"fmt"

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

	err := s.config.SetUser(cmd.name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set", cmd.name)

	return nil
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	currentCmd, ok := c.commandList[cmd.name]
	if ok {
		err := currentCmd(s, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}
