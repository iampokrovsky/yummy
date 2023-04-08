package service

import (
	"context"
	"hw-5/internal/app/menu/model"
)

type Repo interface {
	Create(ctx context.Context, item model.MenuItem) (model.ID, error)
	GetByID(ctx context.Context, id model.ID) (model.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId model.ID) ([]model.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]model.MenuItem, error)
	Update(ctx context.Context, item model.MenuItem) (bool, error)
	Delete(ctx context.Context, id model.ID) (bool, error)
	Restore(ctx context.Context, id model.ID) (bool, error)
}

type Service struct {
	repo Repo
}

// NewService returns a pointer to a new instance of the Service
func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

// Create creates a menu item
func (r *Service) Create(ctx context.Context, item model.MenuItem) (model.ID, error) {
	return r.repo.Create(ctx, item)
}

// GetByID returns a menu item by ID
func (r *Service) GetByID(ctx context.Context, id model.ID) (model.MenuItem, error) {
	return r.GetByID(ctx, id)
}

// ListByRestaurantID returns menu items by restaurant ID
func (r *Service) ListByRestaurantID(ctx context.Context, restId model.ID) ([]model.MenuItem, error) {
	return r.repo.ListByRestaurantID(ctx, restId)
}

// ListByName returns menu items by name match
func (r *Service) ListByName(ctx context.Context, name string) ([]model.MenuItem, error) {
	return r.ListByName(ctx, name)
}

// Update updates a menu item
func (r *Service) Update(ctx context.Context, item model.MenuItem) (bool, error) {
	return r.repo.Update(ctx, item)
}

// Delete removes a menu item by ID
func (r *Service) Delete(ctx context.Context, id model.ID) (bool, error) {
	return r.repo.Delete(ctx, id)
}

// Restore restores a menu item by ID
func (r *Service) Restore(ctx context.Context, id model.ID) (bool, error) {
	return r.repo.Restore(ctx, id)
}
