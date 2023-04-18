package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// RestoreMenuItem restores a menu item.
func (s *CoreService) RestoreMenuItem(ctx context.Context, id menumodel.ID) (bool, error) {
	return s.menuService.Restore(ctx, id)
}
