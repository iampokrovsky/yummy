package postgres

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DBops interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// Get returns a database row matched by args
func (pg PostgresDB) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, pg.pool, dest, query, args...)
}

// Select returns a database rows matched by args
func (pg PostgresDB) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, pg.pool, dest, query, args...)
}

// Exec executes commands and returns nothing
func (pg PostgresDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return pg.pool.Exec(ctx, query, args...)
}

// QueryRow executes command and returns a row
func (pg PostgresDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return pg.pool.QueryRow(ctx, query, args...)
}
