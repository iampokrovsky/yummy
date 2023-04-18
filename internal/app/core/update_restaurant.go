package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// UpdateRestaurant updates a restaurant.
func (s *CoreService) UpdateRestaurant(ctx context.Context, item restmodel.Restaurant) (bool, error) {
	return s.restaurantService.Update(ctx, item)
}
