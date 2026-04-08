package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("Invalid command. Usage: browse <optional_number_of_posts>\nDefault number of posts is 2.")
	}

	limit := int32(2) // Default limit
	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil || parsedLimit <= 0 {
			return fmt.Errorf("Invalid number of posts: %v. Please provide a positive integer.", cmd.args[0])
		}
		limit = int32(parsedLimit)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("Error fetching posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found for your followed feeds.")
		return nil
	}

	fmt.Printf("Showing the latest %d posts from your followed feeds:\n", limit)
	for _, post := range posts {
		fmt.Printf("- %s (Published at: %s)\n  URL: %s\n  Description: %s\n\n", post.Title, post.PublishedAt.Format(time.RFC1123), post.Url, post.Description.String)
	}

	return nil
}
