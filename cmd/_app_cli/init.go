package main

import (
	"context"
	"log"
	"yummy/cmd/_app_cli/commands"
	"yummy/config"
	restrepo "yummy/internal/app/_restaurant/repo"
	restservice "yummy/internal/app/_restaurant/service"
	core "yummy/internal/app/core"
	menurepo "yummy/internal/app/menu/repo"
	menuservice "yummy/internal/app/menu/service"
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

	// Init _restaurant service
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
