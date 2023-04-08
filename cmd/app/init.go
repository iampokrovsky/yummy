package main

import (
	"context"
	"hw-5/config"
	"hw-5/internal/app/menu/model"
	"hw-5/internal/app/menu/repo"
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

	menuRepo := repo.NewPostgresRepo(db)

	//item := model.MenuItem{
	//	RestaurantID: 1,
	//	Name:         "Fresh Baked Bread",
	//	Price:        100000,
	//}
	//
	//_, err = menuRepo.Create(context.TODO(), item)
	//if err != nil {
	//	return
	//}

	//item, _ = menuRepo.GetByID(context.TODO(), 101)
	//fmt.Println(item)

	//items, err := menuRepo.ListByRestaurantID(context.TODO(), 10)
	//fmt.Println(items, err)
	//
	//fmt.Println("")

	//items, err = menuRepo.ListByName(context.TODO(), "Baked")
	//fmt.Println(items, err)

	item2 := model.MenuItem{
		ID:   112,
		Name: "Old",
	}

	_, err = menuRepo.Update(context.TODO(), item2)
	if err != nil {
		return
	}

	_, err = menuRepo.Delete(context.TODO(), 112)
	if err != nil {
		return
	}

	_, err = menuRepo.Restore(context.TODO(), 112)
	if err != nil {
		return
	}

}
