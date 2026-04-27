package main

import (
	"context"
	"errors"
	"fmt"
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
		error := fmt.Sprintf("%s already exists", existing.Name)
		return errors.New(error)
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	new, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	s.config.SetUser(userName)

	fmt.Println(new.Name, "has been registered successfully")

	return nil
}
