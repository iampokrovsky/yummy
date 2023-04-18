package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// CreateMenuItem creates a new menu item.
func (s *CoreService) CreateMenuItem(ctx context.Context, item menumodel.MenuItem) (menumodel.ID, error) {
	return s.menuService.Create(ctx, item)
}
