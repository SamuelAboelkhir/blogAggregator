package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		return errors.New("please provide a feed url")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Colleting feeds every %s\n\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	lastFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("!!!Fetching %s's feed!!!\n", lastFeed.Name)
	fmt.Println("----------------------------------")

	params := database.MarkFeedFetchedParams{
		ID:        lastFeed.ID,
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	_, err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), lastFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
