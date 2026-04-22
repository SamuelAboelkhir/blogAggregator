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
