package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"hw-5/internal/app/menu/model"
	"strconv"
	"strings"
	"time"
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

// NewPostgresRepo returns a pointer to a new instance of the PostgresRepo type
func NewPostgresRepo(db DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

// Create creates a menu item
func (r *PostgresRepo) Create(ctx context.Context, item model.MenuItem) (model.ID, error) {
	var id model.ID

	query := `INSERT INTO menu_items(restaurant_id, name, price) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, item.RestaurantID, item.Name, item.Price).Scan(&id)

	return id, err
}

// GetByID returns a menu item by ID
func (r *PostgresRepo) GetByID(ctx context.Context, id model.ID) (model.MenuItem, error) {
	var item model.MenuItem

	query := `SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at FROM menu_items WHERE id = $1`
	if err := r.db.Get(ctx, &item, query, id); err == sql.ErrNoRows {
		return item, ErrObjectNotFound
	}

	return item, nil
}

// ListByRestaurantID returns menu items by restaurant ID
func (r *PostgresRepo) ListByRestaurantID(ctx context.Context, restId model.ID) ([]model.MenuItem, error) {
	var items []model.MenuItem

	query := `SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at FROM menu_items WHERE restaurant_id = $1`
	err := r.db.Select(ctx, &items, query, restId)

	return items, err
}

// ListByName returns menu items by name match
func (r *PostgresRepo) ListByName(ctx context.Context, name string) ([]model.MenuItem, error) {
	var items []model.MenuItem

	query := `SELECT id, restaurant_id, name, price, created_at, updated_at, deleted_at FROM menu_items WHERE name ILIKE $1`
	err := r.db.Select(ctx, &items, query, "%"+name+"%")

	return items, err
}

// Update updates a menu item
func (r *PostgresRepo) Update(ctx context.Context, item model.MenuItem) (bool, error) {
	var query strings.Builder
	query.WriteString(`UPDATE menu_items SET deleted_at = NULL, updated_at = $1`)
	params := []interface{}{time.Now()}

	if item.Name != "" {
		query.WriteString(`, name = $2`)
		params = append(params, item.Name)
	}

	if item.Price != 0 {
		query.WriteString(`, price = $3`)
		params = append(params, item.Price)
	}

	query.WriteString(` WHERE id = $`)
	params = append(params, item.ID)
	query.WriteString(strconv.Itoa(len(params)))

	result, err := r.db.Exec(ctx, query.String(), params...)

	return result.RowsAffected() > 0, err
}

// Delete removes a menu item by ID
func (r *PostgresRepo) Delete(ctx context.Context, id model.ID) (bool, error) {
	query := `UPDATE menu_items SET deleted_at = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, time.Now(), id)

	return result.RowsAffected() > 0, err
}

// Restore restores a menu item by ID
func (r *PostgresRepo) Restore(ctx context.Context, id model.ID) (bool, error) {
	query := `UPDATE menu_items SET deleted_at = NULL, updated_at = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, time.Now(), id)

	return result.RowsAffected() > 0, err
}
