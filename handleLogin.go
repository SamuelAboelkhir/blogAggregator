package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		return errors.New("please provide a username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatalf("User does not exist")
	}
	if err := s.config.SetUser(user.Name); err != nil {
		log.Fatalf("Error loging in")
	}

	fmt.Println("Login successfull")
	return nil
}
