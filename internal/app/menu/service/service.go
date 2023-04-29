package service

import (
	"context"
	"yummy/internal/app/menu/model"
	"yummy/internal/pkg/api"
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
	api.UnimplementedMenuServiceServer
	repo Repo
}

// NewService returns a pointer to a new instance of the Service
func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

// Create creates a menu item
func (s *Service) Create(ctx context.Context, item model.MenuItem) (model.ID, error) {
	return s.repo.Create(ctx, item)
}

// GetByID returns a menu item by ID
func (s *Service) GetByID(ctx context.Context, id model.ID) (model.MenuItem, error) {
	return s.repo.GetByID(ctx, id)
}

// ListByRestaurantID returns menu items by restaurant ID
func (s *Service) ListByRestaurantID(ctx context.Context, restId model.ID) ([]model.MenuItem, error) {
	return s.repo.ListByRestaurantID(ctx, restId)
}

// ListByName returns menu items by name match
func (s *Service) ListByName(ctx context.Context, name string) ([]model.MenuItem, error) {
	return s.repo.ListByName(ctx, name)
}

// Update updates a menu item
func (s *Service) Update(ctx context.Context, item model.MenuItem) (bool, error) {
	return s.repo.Update(ctx, item)
}

// Delete removes a menu item by ID
func (s *Service) Delete(ctx context.Context, id model.ID) (bool, error) {
	return s.repo.Delete(ctx, id)
}

// Restore restores a menu item by ID
func (s *Service) Restore(ctx context.Context, id model.ID) (bool, error) {
	return s.repo.Restore(ctx, id)
}
