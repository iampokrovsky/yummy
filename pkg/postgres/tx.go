package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// TxDB is an interface for database operations with transactions.
type TxDB interface {
	DB
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}

// BeginTx starts a new transaction.
func (pg PostgresDB) BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error) {
	var txOpts pgx.TxOptions
	if txOptions != nil {
		txOpts = *txOptions
	}
	tx, err := pg.pool.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("can't begin tx: %w", err)
	}
	return &TxWrapper{pgxTx: tx}, nil
}

// Tx is an interface for database transactions.
type Tx interface {
	DBops
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Close(ctx context.Context) error
}

// TxWrapper is a wrapper for pgx.Tx.
type TxWrapper struct {
	pgxTx pgx.Tx
}

// Commit commits the transaction.
func (tx *TxWrapper) Commit(ctx context.Context) error {
	return tx.pgxTx.Commit(ctx)
}

// Rollback rolls back the transaction.
func (tx *TxWrapper) Rollback(ctx context.Context) error {
	return tx.pgxTx.Rollback(ctx)
}

// Close closes the connection, releasing any open resources.
func (tx *TxWrapper) Close(ctx context.Context) error {
	conn := tx.pgxTx.Conn()

	if conn != nil {
		if err := conn.Close(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Get returns a database row matched by args.
func (tx *TxWrapper) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, tx.pgxTx, dest, query, args...)
}

// Select returns a database rows matched by args.
func (tx *TxWrapper) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, tx.pgxTx, dest, query, args...)
}

// Exec executes commands and returns nothing.
func (tx *TxWrapper) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return tx.pgxTx.Exec(ctx, query, args...)
}

// QueryRow executes command and returns a row
func (tx *TxWrapper) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return tx.pgxTx.QueryRow(ctx, query, args...)
}
