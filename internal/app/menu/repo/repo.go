package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
	"yummy/internal/app/menu/model"
	"yummy/pkg/postgres"
)

var (
	ErrNotFound   = errors.New("object not found")
	ErrBuildQuery = errors.New("failed to build query")
)

// DB is an interface for a database.
type DB interface {
	Get(ctx context.Context, dest any, query string, args ...any) error
	Select(ctx context.Context, dest any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (postgres.Tx, error)
	Builder() *squirrel.StatementBuilderType
}

// MenuRepository is a repository for menu items.
type MenuRepository struct {
	db DB
}

// NewMenuRepo creates a new instance of a menu repository.
func NewMenuRepo(db DB) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}

// Create creates new menu items.
func (r *MenuRepository) Create(ctx context.Context, items ...model.MenuItem) ([]uint64, error) {
	// Build a query
	qb := r.db.Builder().
		Insert("menu_items").
		Columns("restaurant_id", "name", "price")
	for _, item := range items {
		qb = qb.Values(item.RestaurantID, item.Name, item.Price)
	}
	qb = qb.Suffix("RETURNING id")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, ErrBuildQuery
	}

	// Execute the query and read the result
	ids := make([]uint64, 0, len(items))
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
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
func (r *MenuRepository) GetByID(ctx context.Context, id uint64) (model.MenuItem, error) {
	items, err := r.ListByID(ctx, id)
	if err != nil {
		return model.MenuItem{}, err
	}
	if len(items) == 0 {
		return model.MenuItem{}, ErrNotFound
	}

	return items[0], nil
}

// ListByID returns menu items by IDs.
func (r *MenuRepository) ListByID(ctx context.Context, ids ...uint64) ([]model.MenuItem, error) {
	var items []model.MenuItem

	query, args, err := r.db.Builder().
		Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.Eq{"id": ids}).ToSql()
	if err != nil {
		return nil, err
	}

	if err = r.db.Select(ctx, &items, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return items, nil
}

// ListByRestaurantID returns menu items by restaurant ID.
func (r *MenuRepository) ListByRestaurantID(ctx context.Context, restId uint64) ([]model.MenuItem, error) {
	var items []model.MenuItem

	query, args, err := r.db.Builder().
		Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.Eq{"restaurant_id": restId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err = r.db.Select(ctx, &items, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrNotFound
	}

	return items, nil
}

// ListByName returns menu items by name match.
func (r *MenuRepository) ListByName(ctx context.Context, name string) ([]model.MenuItem, error) {
	var items []model.MenuItem

	query, args, err := r.db.Builder().
		Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
		From("menu_items").
		Where(squirrel.ILike{"name": "%" + name + "%"}).ToSql()
	if err != nil {
		return nil, ErrBuildQuery
	}

	if err = r.db.Select(ctx, &items, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrNotFound
	}

	return items, err
}

// Update updates menu items.
func (r *MenuRepository) Update(ctx context.Context, items ...model.MenuItem) (bool, error) {
	var rowsAft int64

	// Start transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	for _, item := range items {
		// Build query
		qb := r.db.Builder().
			Update("menu_items").
			Set("deleted_at", nil).
			Set("updated_at", time.Now())
		if item.Name != "" {
			qb = qb.Set("name", item.Name)
		}
		if item.Price != 0 {
			qb = qb.Set("price", item.Price)
		}
		qb = qb.Where(squirrel.Eq{"id": item.ID})
		query, args, err := qb.ToSql()
		if err != nil {
			return false, ErrBuildQuery
		}

		// Exec query
		result, err := tx.Exec(ctx, query, args...)
		if err != nil {
			tx.Rollback(ctx)
			return false, err
		}

		tx.Commit(ctx)

		rowsAft += result.RowsAffected()
	}

	return int64(len(items)) == rowsAft, nil
}

// Delete removes menu items by ID.
func (r *MenuRepository) Delete(ctx context.Context, ids ...uint64) (bool, error) {
	query, args, err := r.db.Builder().
		Update("menu_items").
		Set("deleted_at", time.Now()).
		Where(squirrel.Eq{"id": ids}).ToSql()
	if err != nil {
		return false, err
	}

	result, err := r.db.Exec(ctx, query, args...)

	return int64(len(ids)) == result.RowsAffected(), err
}

// Restore restores a menu item by ID.
func (r *MenuRepository) Restore(ctx context.Context, ids ...uint64) (bool, error) {
	query, args, err := r.db.Builder().Update("menu_items").
		Set("deleted_at", nil).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": ids}).ToSql()
	if err != nil {
		return false, err
	}

	result, err := r.db.Exec(ctx, query, args...)

	return int64(len(ids)) == result.RowsAffected(), err
}
