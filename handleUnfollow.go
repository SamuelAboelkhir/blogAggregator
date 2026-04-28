package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) <= 0 {
		return errors.New("you must specifiy a feed to unfollow")
	}

	feedURL := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	deleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	deleted, err := s.db.DeleteFeedFollow(context.Background(), deleteParams)
	if err != nil {
		return err
	}

	fmt.Println("Feed follow deleted successfully")
	fmt.Println(deleted.ID)

	return nil
}
