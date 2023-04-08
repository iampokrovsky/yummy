package cli

import (
	"context"
	"fmt"
	menu_model "hw-5/internal/app/menu/model"
	"strconv"
	"strings"
	"sync"
)

func (cli *CLI) menuServiceHandler(ctx context.Context, wg *sync.WaitGroup, cmd string) {
	defer wg.Done()

	switch cmd {
	case "create menu item":
		cli.createMenuItem(ctx)
	case "get menu item":
		cli.getMenuItem(ctx)
	case "list menu items":
		cli.listMenuItems(ctx)
	case "update menu item":
		cli.updateMenuItem(ctx)
	case "delete menu item":
		cli.deleteMenuItem(ctx)
	case "restore menu item":
		cli.restoreMenuItem(ctx)
	default:
		return
	}

}

func (cli *CLI) createMenuItem(ctx context.Context) {
	fmt.Println("Enter menu item name:")
	name, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Enter menu item price:")
	priceStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	priceStr = strings.TrimSpace(priceStr)

	var price int64
	if priceStr != "" {
		price, err = strconv.ParseInt(priceStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	newItem := menu_model.MenuItem{
		Name:  name,
		Price: price,
	}

	id, err := cli.menuService.Create(ctx, newItem)
	if err != nil {
		fmt.Printf("Failed to create menu item: %v\n", err)
		return
	}

	fmt.Printf("Menu item created with ID %d.\n", id)
}

func (cli *CLI) getMenuItem(ctx context.Context) {
	fmt.Println("Enter menu item ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	var id int64
	if idStr != "" {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	item, err := cli.menuService.GetByID(ctx, menu_model.ID(id))
	if err != nil {
		fmt.Printf("Failed to get menu item: %v\n", err)
		return
	}

	fmt.Printf("Menu item: %+v\n", item)
}

func (cli *CLI) listMenuItems(ctx context.Context) {
	fmt.Println("Enter restaurant ID:")
	restIdStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	restIdStr = strings.TrimSpace(restIdStr)

	var restId int64
	if restIdStr != "" {
		restId, err = strconv.ParseInt(restIdStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	items, err := cli.menuService.ListByRestaurantID(ctx, menu_model.ID(restId))
	if err != nil {
		fmt.Printf("Failed to list menu items: %v\n", err)
		return
	}

	fmt.Println("Menu items:")
	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}

func (cli *CLI) updateMenuItem(ctx context.Context) {
	fmt.Println("Enter menu item ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	var id int64
	if idStr != "" {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Enter menu item name:")
	name, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Enter menu item price:")
	priceStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	priceStr = strings.TrimSpace(priceStr)

	var price int64
	if priceStr != "" {
		price, err = strconv.ParseInt(priceStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	newItem := menu_model.MenuItem{
		ID:    menu_model.ID(id),
		Name:  name,
		Price: price,
	}

	ok, err := cli.menuService.Update(ctx, newItem)
	if !ok || err != nil {
		fmt.Printf("Failed to update the menu item: %v\n", err)
		return
	}

	fmt.Println("Menu successfully updated.")
}

func (cli *CLI) deleteMenuItem(ctx context.Context) {
	fmt.Println("Enter menu item ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	var id int64
	if idStr != "" {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	ok, err := cli.menuService.Delete(ctx, menu_model.ID(id))
	if !ok || err != nil {
		fmt.Printf("Failed to delete the restaurant: %v\n", err)
		return
	}

	fmt.Println("menu successfully deleted.")
}

func (cli *CLI) restoreMenuItem(ctx context.Context) {
	fmt.Println("Enter menu item ID:")
	idStr, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	idStr = strings.TrimSpace(idStr)

	var id int64
	if idStr != "" {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	ok, err := cli.menuService.Restore(ctx, menu_model.ID(id))
	if !ok || err != nil {
		fmt.Printf("Failed to restore the menu item: %v\n", err)
		return
	}

	fmt.Println("Menu item successfully restored.")
}
