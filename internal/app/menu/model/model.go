package model

import (
	"database/sql"
	"time"
)

type ID uint64
type Money int64

type MenuItem struct {
	ID           ID           `json:"id" db:"id"`
	RestaurantID ID           `json:"restaurantID" db:"restaurant_id"`
	Name         string       `json:"name" db:"name"`
	Price        Money        `json:"price" db:"price"`
	CreatedAt    time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt    sql.NullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deletedAt" db:"deleted_at"`
}
