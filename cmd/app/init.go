package main

import (
	"context"
	"hw-5/cli"
	"hw-5/config"
	menurepo "hw-5/internal/app/menu/repo"
	menuserv "hw-5/internal/app/menu/service"
	restrepo "hw-5/internal/app/restaurant/repo"
	restserv "hw-5/internal/app/restaurant/service"
	"hw-5/pkg/postgres"
	"log"
)

func run(cfg config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := postgres.NewDB(ctx, cfg.DB.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	restaurantRepo := restrepo.NewPostgresRepo(db)
	restaurantService := restserv.NewService(restaurantRepo)

	menuRepo := menurepo.NewPostgresRepo(db)
	menuService := menuserv.NewService(menuRepo)

	console := cli.NewCLI(restaurantService, menuService)
	console.HandleCmd(ctx)
}
