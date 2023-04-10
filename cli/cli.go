package cli

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	menu_model "hw-5/internal/app/menu/model"
	rest_model "hw-5/internal/app/restaurant/model"
	"log"
	"os"
)

var (
	ErrorInvalidParam    = errors.New("invalid param")
	ErrorRepeatedParams  = errors.New("repeated parameters")
	ErrorNotEnoughParams = errors.New("not enough parameters")
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

func (cli *CLI) HandleCmd(ctx context.Context) {
	actionDesc := "Action: create, get, list, update, delete, restore"
	action := flag.String("a", "", actionDesc)
	actionFull := flag.String("action", "", actionDesc)

	targetDesc := "Action target: restaurant or menu"
	target := flag.String("t", "", targetDesc)
	targetFull := flag.String("target", "", targetDesc)

	dataDesc := "Data in JSON format"
	data := flag.String("d", "", dataDesc)
	dataFull := flag.String("data", "", dataDesc)

	flag.Parse()

	targets := []string{"menu", "restaurant"}
	tg, err := validateParam(ctx, targets, *target, *targetFull)
	if err != nil {
		log.Fatal(err)
	}

	actions := []string{"create", "get", "list", "update", "delete", "restore"}
	act, err := validateParam(ctx, actions, *action, *actionFull)
	if err != nil {
		log.Fatal(err)
	}

	dt, err := validateParam(ctx, nil, *data, *dataFull)
	if err != nil {
		log.Fatal(err)
	}

	switch tg {
	case "menu":
		switch act {
		case "create":
			cli.createMenuItem(ctx, dt)
		case "get":
			cli.getMenuItem(ctx, dt)
		case "list":
			cli.listMenuItems(ctx, dt)
		case "update":
			cli.updateMenuItem(ctx, dt)
		case "delete":
			cli.deleteMenuItem(ctx, dt)
		case "restore":
			cli.restoreMenuItem(ctx, dt)
		}
	case "restaurant":
		switch act {
		case "create":
			cli.createRestaurant(ctx, dt)
		case "get":
			cli.getRestaurant(ctx, dt)
		case "list":
			cli.listRestaurants(ctx, dt)
		case "update":
			cli.updateRestaurant(ctx, dt)
		case "delete":
			cli.deleteRestaurant(ctx, dt)
		case "restore":
			cli.restoreRestaurant(ctx, dt)
		}
	}
}

func validateParam(ctx context.Context, validParams []string, param, fullParam string) (string, error) {
	if param == "" && fullParam == "" {
		return "", ErrorNotEnoughParams
	}

	if param != "" && fullParam != "" {
		return "", fmt.Errorf("%v: -%s and --%s", ErrorRepeatedParams, param, fullParam)
	}

	if param == "" && fullParam != "" {
		param = fullParam
	}

	if validParams == nil {
		return param, nil
	}

	var isValid bool

	for _, valid := range validParams {
		if param == valid {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", fmt.Errorf("%v: %s", ErrorInvalidParam, param)
	}

	return param, nil
}
