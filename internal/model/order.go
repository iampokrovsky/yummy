package model

import (
	"database/sql"
	"time"
)

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusReceived              = "received"
	OrderStatusPending               = "pending"
	OrderStatusPreparing             = "preparing"
	OrderStatusShipping              = "shipping"
	OrderStatusCompleted             = "completed"
	OrderStatusCancelled             = "cancelled "
	OrderStatusFailed                = "failed"
)

type OrderMenuItem struct {
	MenuItem
	ID        ID           `db:"id"`
	Amount    int          `db:"amount"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type OrderTracking struct {
	ID         ID           `db:"id"`
	OrderID    ID           `db:"order_id"`
	Status     OrderStatus  `db:"status"`
	StartedAt  time.Time    `db:"started_at"`
	FinishedAt sql.NullTime `db:"finished_at"`
}

// TODO Нужно ли это?
func NewOrderTracking(orderID ID, status OrderStatus) OrderTracking {
	return OrderTracking{
		OrderID:   orderID,
		Status:    status,
		StartedAt: time.Now(),
	}
}

type Order struct {
	ID           ID           `db:"id"`
	ClientID     ID           `db:"client_id"`
	CourierID    ID           `db:"courier_id"`
	RestaurantID ID           `db:"restaurant_id"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
	Items        []OrderMenuItem
	Tracking     []OrderTracking
}
