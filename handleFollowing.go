package main

import (
	"context"
	"fmt"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
