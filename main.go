package main

import (
	"fmt"
	"log"

	"github.com/SamuelAboelkhir/blogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	if err := cfg.SetUser("Samuel"); err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
