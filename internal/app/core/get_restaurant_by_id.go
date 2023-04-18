package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// GetRestaurantByID returns a restaurant by ID.
func (s *CoreService) GetRestaurantByID(ctx context.Context, id restmodel.ID) (restmodel.Restaurant, error) {
	return s.restaurantService.GetByID(ctx, id)
}
