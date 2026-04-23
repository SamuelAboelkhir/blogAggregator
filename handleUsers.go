package main

import (
	"context"
	"fmt"
)

func handleUsers(s *state, cmd command) error {
	if users, err := s.db.GetUsers(context.Background()); err != nil {
		return err
	} else {
		for _, user := range users {
			if user.Name == s.config.CurrentUserName {
				fmt.Println("*", user.Name, "(current)")
				continue
			}
			fmt.Println("*", user.Name)
		}
	}
	return nil
}
