package service

import (
	"context"
	"errors"
	"hw-5/internal/app/restaurant/model"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type Repo interface {
	Create(ctx context.Context, item model.Restaurant) (model.ID, error)
	GetByID(ctx context.Context, id model.ID) (model.Restaurant, error)
	List(ctx context.Context) ([]model.Restaurant, error)
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

// Create creates a restaurant
func (r *Service) Create(ctx context.Context, item model.Restaurant) (model.ID, error) {
	return r.repo.Create(ctx, item)
}

// GetByID returns a restaurant by ID
func (r *Service) GetByID(ctx context.Context, id model.ID) (model.Restaurant, error) {
	return r.repo.GetByID(ctx, id)
}

// List returns a list of all restaurants
func (r *Service) List(ctx context.Context) ([]model.Restaurant, error) {
	return r.repo.List(ctx)
}

// ListByName returns restaurants by name match
func (r *Service) ListByName(ctx context.Context, name string) ([]model.Restaurant, error) {
	return r.repo.ListByName(ctx, name)
}

// ListByCuisine returns menu items by cuisine match
func (r *Service) ListByCuisine(ctx context.Context, cuisine string) ([]model.Restaurant, error) {
	return r.repo.ListByCuisine(ctx, cuisine)
}

// Update updates a restaurant
func (r *Service) Update(ctx context.Context, item model.Restaurant) (bool, error) {
	return r.repo.Update(ctx, item)
}

// Delete removes a restaurant by ID
func (r *Service) Delete(ctx context.Context, id model.ID) (bool, error) {
	return r.repo.Delete(ctx, id)
}

// Restore restores a restaurant item by ID
func (r *Service) Restore(ctx context.Context, id model.ID) (bool, error) {
	return r.repo.Restore(ctx, id)
}
