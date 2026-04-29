package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
	"github.com/google/uuid"
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

func scrapeFeeds(s *state) {
	lastFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
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
		log.Printf("Couldn't mark feed %s fetched: %v", lastFeed.Name, err)
		return
	}

	feed, err := fetchFeed(context.Background(), lastFeed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", lastFeed.Name, err)
		return
	}

	for _, item := range feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		saveParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      lastFeed.ID,
		}

		savedPost, err := s.db.CreatePost(context.Background(), saveParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				fmt.Printf("Post %s, already exists\n", item.Title)
				continue
			}
		}
		fmt.Printf("Post, %s, has been added\n", savedPost.Title)

	}
	fmt.Printf("All posts from %s have been collected, total number of posts %v\n", feed.Channel.Title, len(feed.Channel.Item))
}
