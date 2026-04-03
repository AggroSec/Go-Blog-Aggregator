package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	fmt.Println("Resetting database...")
	return s.db.Reset(context.Background())
}
