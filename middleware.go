package main

import (
	"context"
	"fmt"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name)
		if err != nil {
			return fmt.Errorf("you must be logged in to use this command")
		}
		return handler(s, cmd, user)
	}
}
