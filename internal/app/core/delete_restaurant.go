package core

import (
	"context"
	restmodel "yummy/internal/app/_restaurant/model"
)

// DeleteRestaurant deletes a _restaurant by ID.
func (s *CoreService) DeleteRestaurant(ctx context.Context, id restmodel.ID) (bool, error) {
	return s.restaurantService.Delete(ctx, id)
}
