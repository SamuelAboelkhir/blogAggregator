package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("please provide a name and url")
	}

	currentUser := s.config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		log.Fatal(err)
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		log.Fatal(err)
	}

	feedParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	newFollow, err := s.db.CreateFeedFollow(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully")
	fmt.Println(feed)
	fmt.Println("Feed added to user's feeds")
	fmt.Println(newFollow.FeedName, currentUser)

	return nil
}
