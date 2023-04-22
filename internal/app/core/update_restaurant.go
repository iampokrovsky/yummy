package core

import (
	"context"
	restmodel "yummy/internal/app/_restaurant/model"
)

// UpdateRestaurant updates a _restaurant.
func (s *CoreService) UpdateRestaurant(ctx context.Context, item restmodel.Restaurant) (bool, error) {
	return s.restaurantService.Update(ctx, item)
}
