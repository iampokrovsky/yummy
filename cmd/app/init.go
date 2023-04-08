package main

import (
	"context"
	"hw-5/cli"
	"hw-5/config"
	menu_repo "hw-5/internal/app/menu/repo"
	menu_service "hw-5/internal/app/menu/service"
	rest_repo "hw-5/internal/app/restaurant/repo"
	rest_service "hw-5/internal/app/restaurant/service"
	"hw-5/pkg/postgres"
	"log"
)

func run(cfg config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := postgres.New(ctx, cfg.DB.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	restaurantRepo := rest_repo.NewPostgresRepo(db)
	menuRepo := menu_repo.NewPostgresRepo(db)

	restaurantService := rest_service.NewService(restaurantRepo)
	menuService := menu_service.NewService(menuRepo)

	console := cli.NewCLI(restaurantService, menuService)

	console.Run(ctx)
}
