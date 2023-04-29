package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"yummy/config"
	menu_repo "yummy/internal/app/menu/repo"
	"yummy/internal/pkg/rest"
	"yummy/pkg/httpserver"
	"yummy/test/postgres"
)

func Run(cfg config.Config) {
	// Init layers
	db := postgres.NewPostgresTestDB(context.Background(), cfg.DB.GetDSN())
	menuRepo := menu_repo.NewMenuRepo(db)
	router := rest.NewRouter(menuRepo)
	server := httpserver.New(router, httpserver.Port("8080"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("Got signal: %v, exiting.", s)
	case err := <-server.Notify():
		log.Printf("Got server err signal: %v, exiting.", err)
	}

	// Shutdown
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server shutdown error: %s", err)
	}
}
