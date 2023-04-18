package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// GetMenuItemByID returns a menu item by ID.
func (s *CoreService) GetMenuItemByID(ctx context.Context, id menumodel.ID) (menumodel.MenuItem, error) {
	return s.menuService.GetByID(ctx, id)
}
