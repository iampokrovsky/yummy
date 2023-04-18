package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// ListAllRestaurants returns a list of all restaurants.
func (s *CoreService) ListAllRestaurants(ctx context.Context) ([]restmodel.Restaurant, error) {
	return s.restaurantService.ListAll(ctx)
}
