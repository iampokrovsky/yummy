package postgres

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// DBops is an interface for database operations.
type DBops interface {
	Get(ctx context.Context, dest any, query string, args ...any) error
	Select(ctx context.Context, dest any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// DB is an interface for database.
type DB interface {
	DBops
	Close()
}

// Tx is an interface for database transactions.
type Tx interface {
	DBops
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Close(ctx context.Context) error
}

// TxDB is an interface for database operations with transactions.
type TxDB interface {
	DB
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}
