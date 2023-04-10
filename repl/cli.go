package repl

import (
	"bufio"
	"context"
	"fmt"
	menu_model "hw-5/internal/app/menu/model"
	rest_model "hw-5/internal/app/restaurant/model"
	"os"
	"strings"
	"sync"
)

type RestaurantService interface {
	Create(ctx context.Context, item rest_model.Restaurant) (rest_model.ID, error)
	GetByID(ctx context.Context, id rest_model.ID) (rest_model.Restaurant, error)
	List(ctx context.Context) ([]rest_model.Restaurant, error)
	ListByName(ctx context.Context, name string) ([]rest_model.Restaurant, error)
	ListByCuisine(ctx context.Context, cuisine string) ([]rest_model.Restaurant, error)
	Update(ctx context.Context, item rest_model.Restaurant) (bool, error)
	Delete(ctx context.Context, id rest_model.ID) (bool, error)
	Restore(ctx context.Context, id rest_model.ID) (bool, error)
}

type MenuService interface {
	Create(ctx context.Context, item menu_model.MenuItem) (menu_model.ID, error)
	GetByID(ctx context.Context, id menu_model.ID) (menu_model.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId menu_model.ID) ([]menu_model.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]menu_model.MenuItem, error)
	Update(ctx context.Context, item menu_model.MenuItem) (bool, error)
	Delete(ctx context.Context, id menu_model.ID) (bool, error)
	Restore(ctx context.Context, id menu_model.ID) (bool, error)
}

type CLI struct {
	restaurantService RestaurantService
	menuService       MenuService
	reader            *bufio.Reader
}

func NewCLI(restaurantService RestaurantService, menuService MenuService) *CLI {
	return &CLI{
		restaurantService: restaurantService,
		menuService:       menuService,
		reader:            bufio.NewReader(os.Stdin),
	}
}

func (cli *CLI) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for {
		fmt.Println("Enter command (create, get, list, update, delete, restore) and object (restaurant, menu):")

		cmd, err := cli.reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd = strings.TrimSpace(cmd)

		wg.Add(2)
		go cli.restaurantServiceHandler(ctx, &wg, cmd)
		go cli.menuServiceHandler(ctx, &wg, cmd)
		wg.Wait()
	}
}
