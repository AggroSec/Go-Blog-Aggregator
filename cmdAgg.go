package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	rss "github.com/AggroSec/Go-Blog-Aggregator/internal/RSS"
	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
	"github.com/google/uuid"
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
		dateParsed, err := parseDate(item.PubDate)
		if err != nil {
			fmt.Printf("Warning: could not parse publication date '%s' for item '%s': %v. Using current time instead.\n", item.PubDate, item.Title, err)
			dateParsed = time.Now()
		}
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: dateParsed,
			FeedID:      feed.ID,
		}
		post, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			fmt.Printf("Error creating post for item '%s': %v\n", item.Title, err)
			continue
		}
		fmt.Printf("Created post: %s (ID: %s)\n", post.Title, post.ID)
	}

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,                    // Most common modern format
		"2006-01-02T15:04:05Z07:00",     // RFC3339 with timezone
		"2006-01-02 15:04:05",           // Common DB format
		"2006-01-02",                    // Just date
		"Mon, 02 Jan 2006 15:04:05 MST", // RSS common format
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"January 2, 2006 15:04:05", // Long format
		"2006/01/02 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, dateStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}
