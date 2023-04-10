package model

import (
	"database/sql"
	"time"
)

type ID int64

type MenuItem struct {
	ID           ID           `db:"id" json:"id"`
	RestaurantID ID           `db:"restaurant_id" json:"restaurantID"`
	Name         string       `db:"name" json:"name"`
	Price        int64        `db:"price" json:"price"`
	CreatedAt    time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updatedAt"`
	DeletedAt    sql.NullTime `db:"deleted_at" json:"deletedAt"`
}
