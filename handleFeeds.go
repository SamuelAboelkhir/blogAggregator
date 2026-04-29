package main

import (
	"context"
	"fmt"
)

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("* ID:            %s\n", feed.ID)
		fmt.Printf("* Created:       %v\n", feed.CreatedAt)
		fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)
		fmt.Printf("* LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("* User:          %s\n", user.Name)
	}
	return nil
}
