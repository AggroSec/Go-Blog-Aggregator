package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 || len(cmd.args) > 1 {
		return fmt.Errorf("Invalid command. Usage: follow <feed_url>")
	}

	feed, err := s.db.FeedLookup(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error looking up feed: %v", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %v", err)
	}

	fmt.Printf("Successfully followed feed '%s' (ID: %s) as user '%s' (ID: %s)\n", feedFollow.FeedName, feedFollow.FeedID, feedFollow.UserName, feedFollow.UserID)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Invalid command. Usage: following")
	}

	following, err := s.db.FollowingLookup(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error fetching following list: %v", err)
	}

	if len(following) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Println("You are following these feeds:")
	for _, feed := range following {
		fmt.Printf("- %s (URL: %s)\n", feed.FeedName, feed.FeedUrl)
	}

	return nil
}
