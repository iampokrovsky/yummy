package main

import (
	"hw-5/config"
	"log"
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
