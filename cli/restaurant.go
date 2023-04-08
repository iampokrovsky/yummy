package cli

import (
	"context"
	"fmt"
	rest_model "hw-5/internal/app/restaurant/model"
	"strconv"
	"strings"
	"sync"
)

func (cli *CLI) restaurantServiceHandler(ctx context.Context, wg *sync.WaitGroup, cmd string) {
	defer wg.Done()

	switch cmd {
	case "create restaurant":
		cli.createRestaurant(ctx)
	case "get restaurant":
		cli.getRestaurant(ctx)
	case "list restaurants":
		cli.listRestaurants(ctx)
	case "update restaurant":
		cli.updateRestaurant(ctx)
	case "delete restaurant":
		cli.deleteRestaurant(ctx)
	case "restore restaurant":
		cli.restoreRestaurant(ctx)
	default:
		fmt.Println("Invalid command.")
	}

}

func (cli *CLI) createRestaurant(ctx context.Context) {
	fmt.Println("Enter restaurant name:")
	name, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Enter restaurant address:")
	address, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	address = strings.TrimSpace(address)

	fmt.Println("Enter restaurant cuisine:")
	cuisine, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	cuisine = strings.TrimSpace(cuisine)

	newRestaurant := rest_model.Restaurant{
		Name:    name,
		Address: address,
		Cuisine: rest_model.Cuisine(cuisine),
	}

	id, err := cli.restaurantService.Create(ctx, newRestaurant)
	if err != nil {
		fmt.Printf("Failed to create restaurant: %v\n", err)
		return
	}

	fmt.Printf("Restaurant created with ID %d.\n", id)
}

func (cli *CLI) getRestaurant(ctx context.Context) {
	fmt.Println("Enter restaurant ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)

	restaurant, err := cli.restaurantService.GetByID(ctx, rest_model.ID(id))
	if err != nil {
		fmt.Printf("Failed to get restaurant: %v\n", err)
		return
	}

	fmt.Printf("Restaurant: %+v\n", restaurant)
}

func (cli *CLI) listRestaurants(ctx context.Context) {
	restaurants, err := cli.restaurantService.List(ctx)
	if err != nil {
		fmt.Printf("Failed to list restaurants: %v\n", err)
		return
	}

	fmt.Println("Restaurants:")
	for _, restaurant := range restaurants {
		fmt.Printf("%+v\n", restaurant)
	}
}

func (cli *CLI) updateRestaurant(ctx context.Context) {
	fmt.Println("Enter restaurant ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)

	fmt.Println("Enter restaurant name:")
	name, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Enter restaurant address:")
	address, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	address = strings.TrimSpace(address)

	fmt.Println("Enter restaurant cuisine:")
	cuisine, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	cuisine = strings.TrimSpace(cuisine)

	newRestaurant := rest_model.Restaurant{
		ID:      rest_model.ID(id),
		Name:    name,
		Address: address,
		Cuisine: rest_model.Cuisine(cuisine),
	}

	ok, err := cli.restaurantService.Update(ctx, newRestaurant)
	if !ok || err != nil {
		fmt.Printf("Failed to update the restaurant: %v\n", err)
		return
	}

	fmt.Println("Restaurant successfully updated.")
}

func (cli *CLI) deleteRestaurant(ctx context.Context) {
	fmt.Println("Enter restaurant ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)

	ok, err := cli.restaurantService.Delete(ctx, rest_model.ID(id))
	if !ok || err != nil {
		fmt.Printf("Failed to delete the restaurant: %v\n", err)
		return
	}

	fmt.Println("Restaurant successfully deleted.")
}

func (cli *CLI) restoreRestaurant(ctx context.Context) {
	fmt.Println("Enter restaurant ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)

	ok, err := cli.restaurantService.Restore(ctx, rest_model.ID(id))
	if !ok || err != nil {
		fmt.Printf("Failed to restore the restaurant: %v\n", err)
		return
	}

	fmt.Println("Restaurant successfully restored.")
}
