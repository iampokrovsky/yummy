package main

import (
	"context"
	"log"
	"yummy/cmd/app_cli/commands"
	"yummy/config"
	core "yummy/internal/app/core"
	menurepo "yummy/internal/app/menu/repo"
	menuservice "yummy/internal/app/menu/service"
	restrepo "yummy/internal/app/restaurant/repo"
	restservice "yummy/internal/app/restaurant/service"
	"yummy/pkg/postgres"
)

func run(cfg config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init DB
	db, err := postgres.NewDB(ctx, cfg.DB.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Init restaurant service
	restRepo := restrepo.NewPostgresRepo(db)
	restService := restservice.NewService(restRepo)

	// Init menu service
	menuRepo := menurepo.NewPostgresRepo(db)
	menuService := menuservice.NewService(menuRepo)

	// Init core service
	coreService := core.NewCoreService(restService, menuService)

	// Init CLI
	cli := commands.NewCLI(coreService)
	cli.Execute(ctx)
}
