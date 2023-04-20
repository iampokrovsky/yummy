package repo

import (
	"database/sql"
	"time"
	"yummy/internal/app/menu/model"
)

type IDRow uint64
type MoneyRow int64

type MenuItemRow struct {
	ID           IDRow        `db:"id"`
	RestaurantID IDRow        `db:"restaurant_id"`
	Name         string       `db:"name"`
	Price        MoneyRow     `db:"price"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (row *MenuItemRow) mapFromModel(m model.MenuItem) {
	row.ID = IDRow(m.ID)
	row.RestaurantID = IDRow(m.RestaurantID)
	row.Name = m.Name
	row.Price = MoneyRow(m.Price)
	row.CreatedAt = m.CreatedAt
	row.UpdatedAt = m.UpdatedAt
	row.DeletedAt = m.DeletedAt
}

func (row *MenuItemRow) mapToModel() model.MenuItem {
	return model.MenuItem{
		ID:           model.ID(row.ID),
		RestaurantID: model.ID(row.RestaurantID),
		Name:         row.Name,
		Price:        model.Money(row.Price),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
		DeletedAt:    row.DeletedAt,
	}
}

func mapFromModels(items []model.MenuItem) []MenuItemRow {
	rows := make([]MenuItemRow, len(items))
	for i, item := range items {
		rows[i].mapFromModel(item)
	}
	return rows
}

func mapToModels(rows []MenuItemRow) []model.MenuItem {
	items := make([]model.MenuItem, len(rows))
	for i, row := range rows {
		items[i] = row.mapToModel()
	}
	return items
}
