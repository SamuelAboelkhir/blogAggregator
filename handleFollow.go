package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		log.Fatal("A URL is required")
	}
	url := cmd.Args[0]

	currentUser := s.config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	newFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(newFollow.FeedName, currentUser)
	return nil
}
