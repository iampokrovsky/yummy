package core

import (
	"context"
	restmodel "yummy/internal/app/_restaurant/model"
)

// CreateRestaurant creates a new _restaurant.
func (s *CoreService) CreateRestaurant(ctx context.Context, item restmodel.Restaurant) (restmodel.ID, error) {
	return s.restaurantService.Create(ctx, item)
}
