package postgres

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Get returns a database row matched by args.
func (pg PostgresDB) Get(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Get(ctx, pg.pool, dest, query, args...)
}

// Select returns a database rows matched by args.
func (pg PostgresDB) Select(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Select(ctx, pg.pool, dest, query, args...)
}

// Exec executes commands and returns nothing.
func (pg PostgresDB) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return pg.pool.Exec(ctx, query, args...)
}

// Query executes command and returns rows.
func (pg PostgresDB) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return pg.pool.Query(ctx, query, args...)
}

// QueryRow executes command and returns a row.
func (pg PostgresDB) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return pg.pool.QueryRow(ctx, query, args...)
}

// SendBatch executes a batch of queries and returns the results.
func (pg PostgresDB) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return pg.pool.SendBatch(ctx, batch)
}

// CopyFrom is used to perform high-performance bulk insert into a table.
func (pg PostgresDB) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return pg.pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
}
