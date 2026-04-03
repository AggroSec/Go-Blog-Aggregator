package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: register <username>")
	}

	name := cmd.args[0]

	createdUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), createdUser)
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}
	fmt.Printf("User '%s' created successfully with ID: %s\n", user.Name, user.ID)

	s.conf.SetUser(user.Name)

	return nil
}
