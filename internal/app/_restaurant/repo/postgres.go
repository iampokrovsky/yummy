package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
	"yummy/internal/app/_restaurant/model"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type DB interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

type PostgresRepo struct {
	db DB
}

// NewPostgresRepo returns a pointer to a new instance of the PostgresRepo
func NewPostgresRepo(db DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

// Create creates a _restaurant
func (r *PostgresRepo) Create(ctx context.Context, item model.Restaurant) (model.ID, error) {
	var id model.ID

	query := `INSERT INTO restaurants (name, address, cuisine) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, item.Name, item.Address, item.Cuisine).Scan(&id)

	return id, err
}

// GetByID returns a _restaurant by ID
func (r *PostgresRepo) GetByID(ctx context.Context, id model.ID) (model.Restaurant, error) {
	var item model.Restaurant

	query := `SELECT id, name, address, cuisine, created_at, updated_at, deleted_at FROM restaurants WHERE id = $1`
	if err := r.db.Get(ctx, &item, query, id); err == sql.ErrNoRows {
		return item, ErrObjectNotFound
	}

	return item, nil
}

// ListAll returns a list of all restaurants
func (r *PostgresRepo) ListAll(ctx context.Context) ([]model.Restaurant, error) {
	var items []model.Restaurant

	query := `SELECT id, name, address, cuisine, created_at, updated_at, deleted_at FROM restaurants`
	err := r.db.Select(ctx, &items, query)

	return items, err
}

// ListByName returns restaurants by name match
func (r *PostgresRepo) ListByName(ctx context.Context, name string) ([]model.Restaurant, error) {
	var items []model.Restaurant

	query := `SELECT id, name, address, cuisine, created_at, updated_at, deleted_at FROM restaurants WHERE name ILIKE $1`
	err := r.db.Select(ctx, &items, query, "%"+name+"%")

	return items, err
}

// ListByCuisine returns menu items by cuisine match
func (r *PostgresRepo) ListByCuisine(ctx context.Context, cuisine string) ([]model.Restaurant, error) {
	var items []model.Restaurant

	query := `SELECT id, name, address, cuisine, created_at, updated_at, deleted_at FROM restaurants WHERE cuisine ILIKE $1`
	err := r.db.Select(ctx, &items, query, "%"+cuisine+"%")

	return items, err
}

// Update updates a _restaurant
func (r *PostgresRepo) Update(ctx context.Context, item model.Restaurant) (bool, error) {
	var query strings.Builder
	query.WriteString(`UPDATE restaurants SET deleted_at = NULL, updated_at = $1`)
	params := []interface{}{time.Now()}

	if item.Name != "" {
		query.WriteString(`, name = $2`)
		params = append(params, item.Name)
	}

	if item.Address != "" {
		query.WriteString(`, address = $3`)
		params = append(params, item.Address)
	}

	if item.Cuisine != "" {
		query.WriteString(`, cuisine = $4`)
		params = append(params, item.Cuisine)
	}

	query.WriteString(` WHERE id = $`)
	params = append(params, item.ID)
	query.WriteString(strconv.Itoa(len(params)))

	result, err := r.db.Exec(ctx, query.String(), params...)

	return result.RowsAffected() > 0, err
}

// Delete removes a _restaurant by ID
func (r *PostgresRepo) Delete(ctx context.Context, id model.ID) (bool, error) {
	query := `UPDATE restaurants SET deleted_at = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, time.Now(), id)

	return result.RowsAffected() > 0, err
}

// Restore restores a _restaurant item by ID
func (r *PostgresRepo) Restore(ctx context.Context, id model.ID) (bool, error) {
	query := `UPDATE restaurants SET deleted_at = NULL, updated_at = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, time.Now(), id)

	return result.RowsAffected() > 0, err
}
