package main

import (
	"context"

	"github.com/SamuelAboelkhir/blogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser := s.config.CurrentUserName
		user, err := s.db.GetUser(context.Background(), currentUser)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
