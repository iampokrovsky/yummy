package core

import (
	"context"
	restaurantmodel "yummy/internal/app/_restaurant/model"
)

// RestoreRestaurant restores a _restaurant.
func (s *CoreService) RestoreRestaurant(ctx context.Context, id restaurantmodel.ID) (bool, error) {
	return s.restaurantService.Restore(ctx, id)
}
