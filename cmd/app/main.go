package main

import (
	"hw-5/config"
	"hw-5/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
