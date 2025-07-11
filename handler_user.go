package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	err := s.cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User switched successfully!")
	return nil
}
