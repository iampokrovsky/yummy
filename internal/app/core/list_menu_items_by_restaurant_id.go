package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// ListMenuItemsByRestaurantID returns a list of menu items filtered by restaurant ID.
func (s *CoreService) ListMenuItemsByRestaurantID(ctx context.Context, restId menumodel.ID) ([]menumodel.MenuItem, error) {
	return s.menuService.ListByRestaurantID(ctx, restId)
}
