package repo

import (
	"context"
	"yummy/internal/app/menu/model"
)

type MenuRepo interface {
	Create(ctx context.Context, items ...model.MenuItem) ([]uint64, error)
	GetByID(ctx context.Context, id uint64) (model.MenuItem, error)
	ListByID(ctx context.Context, ids ...uint64) ([]model.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId uint64) ([]model.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]model.MenuItem, error)
	Update(ctx context.Context, items ...model.MenuItem) (bool, error)
	Delete(ctx context.Context, ids ...uint64) (bool, error)
	Restore(ctx context.Context, ids ...uint64) (bool, error)
}
