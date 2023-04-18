package core

import (
	"context"
	menumodel "yummy/internal/app/menu/model"
)

// DeleteMenuItem deletes a menu item by ID.
func (s *CoreService) DeleteMenuItem(ctx context.Context, id menumodel.ID) (bool, error) {
	return s.menuService.Delete(ctx, id)
}
