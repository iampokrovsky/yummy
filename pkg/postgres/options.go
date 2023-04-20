package postgres

import "time"

// Option is a function that configures the database.
type Option func(*PostgresDB)

// MaxPoolSize specifies the maximum number of connections in the pool.
func MaxPoolSize(size int) Option {
	return func(c *PostgresDB) {
		c.maxPoolSize = size
	}
}

// ConnAttempts specifies the number of connection attempts.
func ConnAttempts(attempts int) Option {
	return func(c *PostgresDB) {
		c.connAttempts = attempts
	}
}

// ConnTimeout specifies the timeout for connecting to the database.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *PostgresDB) {
		c.connTimeout = timeout
	}
}
