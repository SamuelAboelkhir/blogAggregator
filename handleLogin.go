package main

import (
	"context"
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) <= 0 {
		return errors.New("please provide a username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("Login successfull")
	return nil
}
