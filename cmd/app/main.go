package main

import (
	"log"
	"yummy/config"
)

func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	run(cfg)
}
