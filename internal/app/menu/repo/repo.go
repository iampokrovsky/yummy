package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrBuildQuery     = errors.New("failed to build query")
)

// DBops is an interface for database operations.
type DB interface {
	Get(ctx context.Context, dest any, query string, args ...any) error
	Select(ctx context.Context, dest any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Builder() *squirrel.StatementBuilderType
	Close()
}

// MenuRepo is a repository for menu items.
type MenuRepo struct {
	db DB
}

// NewMenuRepo creates a new instance of a menu repository.
func NewMenuRepo(db DB) *MenuRepo {
	return &MenuRepo{
		db: db,
	}
}

// Create creates new menu items.
func (r *MenuRepo) Create(ctx context.Context, items []MenuItemRow) ([]uint64, error) {
	// Build a query
	qb := r.db.Builder().
		Insert("menu_items").
		Columns("restaurant_id", "name", "price")
	for _, item := range items {
		qb = qb.Values(item.RestaurantID, item.Name, item.Price)
	}
	qb.Suffix(`RETURNING id`)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, ErrBuildQuery
	}

	// Execute the query and read the result
	ids := make([]uint64, 0, len(items))
	rows, err := r.db.Query(ctx, query, args...)
	if err == sql.ErrNoRows {
		return nil, ErrObjectNotFound
	}
	for rows.Next() {
		var id uint64
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// GetByID returns a menu item by ID.
func (r *MenuRepo) GetByID(ctx context.Context, id uint64) (MenuItemRow, error) {
	var item MenuItemRow

	query, args, err := r.db.Builder().Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return item, ErrBuildQuery
	}

	fmt.Println(query, args)

	if err = r.db.Get(ctx, &item, query, args...); err == sql.ErrNoRows {
		return item, ErrObjectNotFound
	}
	if err != nil {
		return item, err
	}

	return item, nil
}

// ListByRestaurantID returns menu items by _restaurant ID.
func (r *MenuRepo) ListByRestaurantID(ctx context.Context, restId uint64) ([]MenuItemRow, error) {
	var items []MenuItemRow

	query, args, err := r.db.Builder().
		Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.Eq{"restaurant_id": restId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err = r.db.Select(ctx, &items, query, args...); err == sql.ErrNoRows {
		return nil, ErrObjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return items, nil
}

// ListByName returns menu items by name match.
func (r *MenuRepo) ListByName(ctx context.Context, name string) ([]MenuItemRow, error) {
	var items []MenuItemRow

	query, args, err := r.db.Builder().Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.ILike{"name": "%" + name + "%"}).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(ctx, &items, query, args...); err != nil {
		return nil, err
	}

	return items, err
}

// Update updates menu items.
func (r *MenuRepo) Update(ctx context.Context, items []MenuItemRow) (bool, error) {
	qb := squirrel.Update("menu_items").
		Set("deleted_at", nil).
		Set("updated_at", time.Now())

	for _, item := range items {
		if item.Name != "" {
			qb = qb.Set("name", item.Name)
		}

		if item.Price != 0 {
			qb = qb.Set("price", item.Price)
		}

		qb = qb.Where(squirrel.Eq{"id": item.ID})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	result, err := r.db.Exec(ctx, query, args...)

	return result.RowsAffected() > 0, err
}

// Delete removes a menu item by ID.
func (r *MenuRepo) Delete(ctx context.Context, ids []uint64) (bool, error) {
	query, args, err := r.db.Builder().Update("menu_items").
		Set("deleted_at", time.Now()).
		Where(squirrel.Eq{"id": ids}).
		ToSql()
	if err != nil {
		return false, err
	}

	result, err := r.db.Exec(ctx, query, args...)

	return result.RowsAffected() > 0, err
}

// Restore restores a menu item by ID.
func (r *MenuRepo) Restore(ctx context.Context, ids []uint64) (bool, error) {
	query, args, err := r.db.Builder().Update("menu_items").
		Set("deleted_at", nil).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": ids}).
		ToSql()
	if err != nil {
		return false, err
	}

	result, err := r.db.Exec(ctx, query, args...)

	return result.RowsAffected() > 0, err
}
