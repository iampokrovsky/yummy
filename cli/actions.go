package cli

import (
	"context"
	"encoding/json"
	"fmt"
	menumodel "hw-5/internal/app/menu/model"
	restmodel "hw-5/internal/app/restaurant/model"
	"log"
)

func (cli *CLI) createMenuItem(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	id, err := cli.menuService.Create(ctx, itemRaw)
	if err != nil {
		log.Fatalf("failed to create menu item: %v\n", err)
	}

	fmt.Printf("Menu item created with ID %d.\n", id)
}

func (cli *CLI) getMenuItem(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	item, err := cli.menuService.GetByID(context.Background(), 101)
	if err != nil {
		log.Fatalf("failed to get menu item: %v\n", err)
	}

	fmt.Printf("Menu item: %+v\n", item)
}

func (cli *CLI) listMenuItems(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	items, err := cli.menuService.ListByRestaurantID(ctx, itemRaw.RestaurantID)
	if err != nil {
		log.Fatalf("failed to list menu items: %v\n", err)
	}

	fmt.Println("Menu items:")
	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}

func (cli *CLI) updateMenuItem(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.menuService.Update(ctx, itemRaw)
	if !ok || err != nil {
		log.Fatalf("failed to update the menu item: %v\n", err)
	}

	fmt.Println("Menu successfully updated.")
}

func (cli *CLI) deleteMenuItem(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.menuService.Delete(ctx, itemRaw.ID)
	if !ok || err != nil {
		log.Fatalf("failed to delete the restaurant: %v\n", err)
	}

	fmt.Println("Menu successfully deleted.")
}

func (cli *CLI) restoreMenuItem(ctx context.Context, data string) {
	var itemRaw menumodel.MenuItem
	if err := json.Unmarshal([]byte(data), &itemRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.menuService.Restore(ctx, itemRaw.ID)
	if !ok || err != nil {
		log.Fatalf("failed to restore the menu item: %v\n", err)
	}

	fmt.Println("Menu item successfully restored.")
}

func (cli *CLI) createRestaurant(ctx context.Context, data string) {
	var restRaw restmodel.Restaurant
	if err := json.Unmarshal([]byte(data), &restRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	id, err := cli.restaurantService.Create(ctx, restRaw)
	if err != nil {
		log.Fatalf("failed to create restaurant: %v\n", err)
	}

	fmt.Printf("Restaurant created with ID %d.\n", id)
}

func (cli *CLI) getRestaurant(ctx context.Context, data string) {
	var restRaw restmodel.Restaurant
	if err := json.Unmarshal([]byte(data), &restRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	restaurant, err := cli.restaurantService.GetByID(ctx, restRaw.ID)
	if err != nil {
		log.Fatalf("failed to get restaurant: %v\n", err)
		return
	}

	fmt.Printf("Restaurant: %+v\n", restaurant)
}

func (cli *CLI) listRestaurants(ctx context.Context, data string) {
	restaurants, err := cli.restaurantService.List(ctx)
	if err != nil {
		log.Fatalf("failed to list restaurants: %v\n", err)
	}

	fmt.Println("Restaurants:")
	for _, restaurant := range restaurants {
		fmt.Printf("%+v\n", restaurant)
	}
}

func (cli *CLI) updateRestaurant(ctx context.Context, data string) {
	var restRaw restmodel.Restaurant
	if err := json.Unmarshal([]byte(data), &restRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.restaurantService.Update(ctx, restRaw)
	if !ok || err != nil {
		log.Fatalf("failed to update the restaurant: %v\n", err)
	}

	fmt.Println("Restaurant successfully updated.")
}

func (cli *CLI) deleteRestaurant(ctx context.Context, data string) {
	var restRaw restmodel.Restaurant
	if err := json.Unmarshal([]byte(data), &restRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.restaurantService.Delete(ctx, restRaw.ID)
	if !ok || err != nil {
		log.Fatalf("failed to delete the restaurant: %v\n", err)
	}

	fmt.Println("Restaurant successfully deleted.")
}

func (cli *CLI) restoreRestaurant(ctx context.Context, data string) {
	var restRaw restmodel.Restaurant
	if err := json.Unmarshal([]byte(data), &restRaw); err != nil {
		log.Fatal("error parsing JSON:", err)
	}

	ok, err := cli.restaurantService.Restore(ctx, restRaw.ID)
	if !ok || err != nil {
		log.Fatalf("failed to restore the restaurant: %v\n", err)
	}

	fmt.Println("Restaurant successfully restored.")
}
