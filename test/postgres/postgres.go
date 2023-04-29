package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
	"sync"
	"testing"
	"yummy/pkg/postgres"
)

// PostgresTestDB is a wrapper around PostgresDB that provides test helpers.
type PostgresTestDB struct {
	sync.Mutex
	*postgres.PostgresDB
}

// NewPostgresTestDB creates a new PostgresTestDB.
func NewPostgresTestDB(ctx context.Context, dsn string) *PostgresTestDB {
	db, err := postgres.NewDB(ctx, dsn)
	if err != nil {
		panic(err)
	}

	return &PostgresTestDB{
		PostgresDB: db,
	}
}

// SetUp prepares the database for testing.
func (pg *PostgresTestDB) SetUp(ctx context.Context, t *testing.T) {
	t.Helper()
	pg.Lock()
	pg.Truncate(ctx)
}

// TearDown cleans up the database after testing.
func (pg *PostgresTestDB) TearDown(ctx context.Context) {
	defer pg.Unlock()
	pg.Truncate(ctx)
}

// Truncate truncates all tables in the database.
func (pg *PostgresTestDB) Truncate(ctx context.Context) {
	var tables []string

	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' AND table_name != 'goose_db_version';`
	if err := pg.Select(ctx, &tables, query); err != nil {
		panic(err)
	}
	if len(tables) == 0 {
		panic(errors.New("no tables found"))
	}

	batch := pgx.Batch{}

	batch.Queue(fmt.Sprintf(`TRUNCATE TABLE %s`, strings.Join(tables, ", ")))
	for _, table := range tables {
		batch.Queue(fmt.Sprintf(`ALTER SEQUENCE %s_id_seq RESTART WITH 1`, table))
	}

	if err := pg.SendBatch(ctx, &batch).Close(); err != nil {
		panic(err)
	}
}
