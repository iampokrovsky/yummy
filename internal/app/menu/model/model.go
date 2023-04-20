package model

import (
	"database/sql"
	"time"
)

type ID uint64
type Money int64

type MenuItem struct {
	ID           ID           `json:"id"`
	RestaurantID ID           `json:"restaurantID"`
	Name         string       `json:"name"`
	Price        Money        `json:"price"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    sql.NullTime `json:"updatedAt"`
	DeletedAt    sql.NullTime `json:"deletedAt"`
}
