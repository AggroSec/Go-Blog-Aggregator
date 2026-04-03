package main

import "fmt"

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 || len(cmd.args) > 1 {
		return fmt.Errorf("Invalid command. Usage: follow <feed_url>")
	}

	//make query to get feed by url

	return nil
}
