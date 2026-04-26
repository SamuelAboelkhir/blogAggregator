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

	userName := s.config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), userName)
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
	result, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Feed created successfully")
	fmt.Println(result)

	return nil
}
