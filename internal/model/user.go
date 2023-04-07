package model

import (
	"database/sql"
	"time"
)

type Transport string

const (
	TransportBicycle Transport = "bicycle"
	TransportCar               = "car"
	TransportScooter           = "scooter"
)

type User struct {
	ID          ID           `db:"id"`
	Name        string       `db:"name"`
	Email       string       `db:"email"`
	PhoneNumber string       `db:"phone_number"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type Client struct {
	User
	ID      ID     `db:"id"`
	Address string `db:"address"`
}

type Courier struct {
	User
	ID        ID        `db:"id"`
	Transport Transport `db:"transport"`
}
