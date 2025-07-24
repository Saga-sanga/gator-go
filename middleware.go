package main

import (
	"context"
	"fmt"

	"github.com/saga-sanga/gator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error retrieving user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
