package Database

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO Посмотреть, какие методы могут быть нужны
type DB interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	GetPool(ctx context.Context) *pgxpool.Pool
}

type Database struct {
	pool *pgxpool.Pool
}

func newDatabase(pool *pgxpool.Pool) *Database {
	return &Database{pool: pool}
}

// TODO Использовать контекст
func (db Database) GetPool(_ context.Context) *pgxpool.Pool {
	return db.pool
}

func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.pool, dest, query, args...)
}

func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.pool, dest, query, args...)
}

func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, query, args...)
}

func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.pool.QueryRow(ctx, query, args...)
}