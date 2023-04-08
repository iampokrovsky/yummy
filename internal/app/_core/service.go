package _core

//
//import (
//	"context"
//	menu_model "hw-5/internal/app/menu/model"
//	rest_model "hw-5/internal/app/restaurant/model"
//)
//

//type RestaurantService interface {
//	Create(ctx context.Context, item rest_model.Restaurant) (rest_model.ID, error)
//	GetByID(ctx context.Context, id rest_model.ID) (rest_model.Restaurant, error)
//	List(ctx context.Context) ([]rest_model.Restaurant, error)
//	ListByName(ctx context.Context, name string) ([]rest_model.Restaurant, error)
//	ListByCuisine(ctx context.Context, cuisine string) ([]rest_model.Restaurant, error)
//	Update(ctx context.Context, item rest_model.Restaurant) (bool, error)
//	Delete(ctx context.Context, id rest_model.ID) (bool, error)
//	Restore(ctx context.Context, id rest_model.ID) (bool, error)
//}
//
//type MenuService interface {
//	Create(ctx context.Context, item menu_model.MenuItem) (menu_model.ID, error)
//	GetByID(ctx context.Context, id menu_model.ID) (menu_model.MenuItem, error)
//	ListByRestaurantID(ctx context.Context, restId menu_model.ID) ([]menu_model.MenuItem, error)
//	ListByName(ctx context.Context, name string) ([]menu_model.MenuItem, error)
//	Update(ctx context.Context, item menu_model.MenuItem) (bool, error)
//	Delete(ctx context.Context, id menu_model.ID) (bool, error)
//	Restore(ctx context.Context, id menu_model.ID) (bool, error)
//}

//
//type Service struct {
//	restaurantService RestaurantService
//	menuService       MenuService
//}
//
//func NewCoreService(restaurantService RestaurantService, menuService MenuService) *Service {
//	return &Service{
//		restaurantService: restaurantService,
//		menuService:       menuService,
//	}
//}
//
//func (s *Service) name() {
//
//}
//
//CreateRestaurant(ctx context.Context, item rest_model.Restaurant) (rest_model.ID, error)
//GetRestaurantByID(ctx context.Context, id rest_model.ID) (rest_model.Restaurant, error)
//ListRestaurants(ctx context.Context) ([]rest_model.Restaurant, error)
//ListRestaurantsByName(ctx context.Context, name string) ([]rest_model.Restaurant, error)
//ListRestaurantsByCuisine(ctx context.Context, cuisine string) ([]rest_model.Restaurant, error)
//UpdateRestaurant(ctx context.Context, item rest_model.Restaurant) (bool, error)
//DeleteRestaurant(ctx context.Context, id rest_model.ID) (bool, error)
//RestoreRestaurant(ctx context.Context, id rest_model.ID) (bool, error)
//
//Create(ctx context.Context, item menu_model.MenuItem) (menu_model.ID, error)
//GetByID(ctx context.Context, id menu_model.ID) (menu_model.MenuItem, error)
//ListByRestaurantID(ctx context.Context, restId menu_model.ID) ([]menu_model.MenuItem, error)
//ListByName(ctx context.Context, name string) ([]menu_model.MenuItem, error)
//Update(ctx context.Context, item menu_model.MenuItem) (bool, error)
//Delete(ctx context.Context, id menu_model.ID) (bool, error)
//Restore(ctx context.Context, id menu_model.ID) (bool, error)
