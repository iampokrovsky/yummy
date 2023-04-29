package service

import (
	"context"
	"yummy/internal/app/_restaurant/model"
)

type Repo interface {
	Create(ctx context.Context, item model.Restaurant) (model.ID, error)
	GetByID(ctx context.Context, id model.ID) (model.Restaurant, error)
	ListAll(ctx context.Context) ([]model.Restaurant, error)
	ListByName(ctx context.Context, name string) ([]model.Restaurant, error)
	ListByCuisine(ctx context.Context, cuisine string) ([]model.Restaurant, error)
	Update(ctx context.Context, item model.Restaurant) (bool, error)
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

// Create creates a _restaurant
func (s *Service) Create(ctx context.Context, item model.Restaurant) (model.ID, error) {
	return s.repo.Create(ctx, item)
}

// GetByID returns a _restaurant by ID
func (s *Service) GetByID(ctx context.Context, id model.ID) (model.Restaurant, error) {
	return s.repo.GetByID(ctx, id)
}

// List returns a list of all restaurants
func (s *Service) ListAll(ctx context.Context) ([]model.Restaurant, error) {
	return s.repo.ListAll(ctx)
}

// ListByName returns restaurants by name match
func (s *Service) ListByName(ctx context.Context, name string) ([]model.Restaurant, error) {
	return s.repo.ListByName(ctx, name)
}

// ListByCuisine returns menu items by cuisine match
func (s *Service) ListByCuisine(ctx context.Context, cuisine string) ([]model.Restaurant, error) {
	return s.repo.ListByCuisine(ctx, cuisine)
}

// Update updates a _restaurant
func (s *Service) Update(ctx context.Context, item model.Restaurant) (bool, error) {
	return s.repo.Update(ctx, item)
}

// Delete removes a _restaurant by ID
func (s *Service) Delete(ctx context.Context, id model.ID) (bool, error) {
	return s.repo.Delete(ctx, id)
}

// Restore restores a _restaurant item by ID
func (s *Service) Restore(ctx context.Context, id model.ID) (bool, error) {
	return s.repo.Restore(ctx, id)
}
