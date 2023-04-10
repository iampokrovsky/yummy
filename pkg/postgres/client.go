package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// New returns pointer to instance of Database
func NewDB(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return newDatabase(pool), nil
}
