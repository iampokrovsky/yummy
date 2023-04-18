package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// UpdateMenuItem updates a menu item.
func (s *CoreService) UpdateMenuItem(ctx context.Context, item menumodel.MenuItem) (bool, error) {
	return s.menuService.Update(ctx, item)
}
