package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

// PostgresDB is a wrapper for pgxpool.Pool.
type PostgresDB struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	builder      *squirrel.StatementBuilderType
	pool         *pgxpool.Pool
}

// NewDB allocates and returns a new PostgresDB.
func NewDB(ctx context.Context, dsn string, opts ...Option) (*PostgresDB, error) {
	pg := &PostgresDB{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pg.builder = &psql

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	for pg.connAttempts > 0 {
		pg.pool, err = pgxpool.ConnectConfig(ctx, poolConfig)
		if err == nil {
			break
		}

		// TODO Add logger
		log.Printf("PostgresDB is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (pg PostgresDB) Builder() *squirrel.StatementBuilderType {
	return pg.builder
}

// Close closes the connection pool, releasing any open resources.
func (pg PostgresDB) Close() {
	if pg.pool != nil {
		pg.pool.Close()
	}
}
