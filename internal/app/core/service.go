package core

import (
	"context"
	restmodel "yummy/internal/app/_restaurant/model"
	menumodel "yummy/internal/app/menu/model"
)

type RestaurantService interface {
	Create(ctx context.Context, item restmodel.Restaurant) (restmodel.ID, error)
	GetByID(ctx context.Context, id restmodel.ID) (restmodel.Restaurant, error)
	ListAll(ctx context.Context) ([]restmodel.Restaurant, error)
	ListByName(ctx context.Context, name string) ([]restmodel.Restaurant, error)
	ListByCuisine(ctx context.Context, cuisine string) ([]restmodel.Restaurant, error)
	Update(ctx context.Context, item restmodel.Restaurant) (bool, error)
	Delete(ctx context.Context, id restmodel.ID) (bool, error)
	Restore(ctx context.Context, id restmodel.ID) (bool, error)
}

type MenuService interface {
	Create(ctx context.Context, item menumodel.MenuItem) (menumodel.ID, error)
	GetByID(ctx context.Context, id menumodel.ID) (menumodel.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId menumodel.ID) ([]menumodel.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]menumodel.MenuItem, error)
	Update(ctx context.Context, item menumodel.MenuItem) (bool, error)
	Delete(ctx context.Context, id menumodel.ID) (bool, error)
	Restore(ctx context.Context, id menumodel.ID) (bool, error)
}

type CoreService struct {
	restaurantService RestaurantService
	menuService       MenuService
}

// NewCoreService creates a new CoreService object and returns the pointer to it.
func NewCoreService(restaurantService RestaurantService, menuService MenuService) *CoreService {
	service := &CoreService{
		restaurantService: restaurantService,
		menuService:       menuService,
	}
	return service
}
