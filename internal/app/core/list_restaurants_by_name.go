package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// ListRestaurantsByName returns a list of restaurants filtered by name.
func (s *CoreService) ListRestaurantsByName(ctx context.Context, name string) ([]restmodel.Restaurant, error) {
	return s.restaurantService.ListByName(ctx, name)
}
