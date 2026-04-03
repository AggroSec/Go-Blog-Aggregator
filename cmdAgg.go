package main

import (
	"context"
	"fmt"

	rss "github.com/AggroSec/Go-Blog-Aggregator/internal/RSS"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%#v", rssFeed)

	return nil
}
