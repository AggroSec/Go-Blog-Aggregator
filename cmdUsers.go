package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, c command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Gator Users:")

	for _, user := range users {
		if user.Name == s.conf.Current_user_name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
