package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// ListRestaurantsByCuisine returns a list of restaurants filtered by cuisine.
func (s *CoreService) ListRestaurantsByCuisine(ctx context.Context, cuisine string) ([]restmodel.Restaurant, error) {
	return s.restaurantService.ListByCuisine(ctx, cuisine)
}
