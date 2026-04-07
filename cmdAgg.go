package main

import (
	"context"
	"fmt"
	"time"

	rss "github.com/AggroSec/Go-Blog-Aggregator/internal/RSS"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Invalid command. Usage: agg <time_between_reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration format: %v", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeErr := scrapeFeeds(s)
		if scrapeErr != nil {
			return scrapeErr
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %v", err)
	}

	fmt.Printf("Fetching feed: %s (URL: %s)\n", feed.Name, feed.Url)

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}

	fmt.Printf("Successfully fetched feed '%s' with %d items.\n", rssFeed.Channel.Title, len(rssFeed.Channel.Item))

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("- %s (Published: %s)\n", item.Title, item.PubDate)
	}

	return nil
}
