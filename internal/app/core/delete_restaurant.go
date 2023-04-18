package core

import (
	"context"
	restmodel "yummy/internal/app/restaurant/model"
)

// DeleteRestaurant deletes a restaurant by ID.
func (s *CoreService) DeleteRestaurant(ctx context.Context, id restmodel.ID) (bool, error) {
	return s.restaurantService.Delete(ctx, id)
}
