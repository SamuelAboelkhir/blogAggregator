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

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		return errors.New("please provide a username")
	}
	userName := cmd.Args[0]

	existing, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		log.Fatalf(existing.Name, "already exists")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	new, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	s.config.SetUser(userName)

	fmt.Println(new.Name, "has been registered successfully")

	return nil
}
