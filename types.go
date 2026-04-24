package main

import (
	"github.com/SamuelAboelkhir/blogAggregator/internal/config"
	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
