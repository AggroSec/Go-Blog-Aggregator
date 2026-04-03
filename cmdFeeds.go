package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.FeedList(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s, URL: %s, Created by: %s\n", feed.Name, feed.Url, feed.Username)
	}

	return nil
}
