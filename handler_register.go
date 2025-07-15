package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/saga-sanga/gator-go/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]

	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User created successfully!")
	fmt.Printf("User: %v", user)

	return nil
}
