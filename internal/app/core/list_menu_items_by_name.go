package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// ListMenuItemsByName returns a list of menu items filtered by name.
func (s *CoreService) ListMenuItemsByName(ctx context.Context, name string) ([]menumodel.MenuItem, error) {
	return s.menuService.ListByName(ctx, name)
}
