package main

import (
	"errors"
)

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	currentCmd, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")

	}
	return currentCmd(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
