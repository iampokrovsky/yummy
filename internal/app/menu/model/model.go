package model

import (
	"database/sql"
	"time"
)

type ID int64

type MenuItem struct {
	ID           ID           `db:"id" `
	RestaurantID ID           `db:"restaurant_id"`
	Name         string       `db:"name"`
	Price        int64        `db:"price"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}
