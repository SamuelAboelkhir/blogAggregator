package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/SamuelAboelkhir/blogAggregator/internal/config"
	"github.com/SamuelAboelkhir/blogAggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)

	dbQueries := database.New(db)

	s := &state{
		db:     dbQueries,
		config: &cfg,
	}

	c := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)

	args := os.Args

	if len(args) <= 2 {
		log.Fatalf("Not enough arguments")
	}

	cmdName := args[1]
	arguments := args[2:]

	command := command{
		Name: cmdName,
		Args: arguments,
	}
	err = c.run(s, command)

	if err != nil {
		log.Fatal(err)
	}
}
