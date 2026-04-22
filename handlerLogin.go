package main

import (
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		return errors.New("Please provide a username")
	}

	if err := s.config.SetUser(cmd.Args[0]); err != nil {
		log.Fatalf("Error setting username")
	}

	fmt.Println("User has been set successfully")
	return nil
}
