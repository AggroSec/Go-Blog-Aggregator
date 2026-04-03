package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if s.conf.Current_user_name == "" {
		return fmt.Errorf("you must be logged in to add a feed")
	}
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: agg addfeed <feed name> <feed url>")
	}

	user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name)
	if err != nil {
		return fmt.Errorf("error fetching user from database: %v", err)
	}

	feed_name := cmd.args[0]
	feed_url := cmd.args[1]

	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feed_name,
		Url:       feed_url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding feed to database: %v", err)
	}

	fmt.Printf("Feed added successfully! ID: %s, Name: %s, URL: %s\n", feed.ID, feed.Name, feed.Url)
	return nil
}
