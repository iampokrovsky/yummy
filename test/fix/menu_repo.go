package fix

import (
	"database/sql"
	"time"
	menurepo "yummy/internal/app/menu/repo"
)

type MenuItemRowBuilder struct {
	instance *menurepo.MenuItemRow
}

func MenuItemRow() *MenuItemRowBuilder {
	return &MenuItemRowBuilder{
		instance: &menurepo.MenuItemRow{},
	}
}

func (b *MenuItemRowBuilder) ID(id uint64) *MenuItemRowBuilder {
	b.instance.ID = menurepo.IDRow(id)
	return b
}

func (b *MenuItemRowBuilder) RestaurantID(id uint64) *MenuItemRowBuilder {
	b.instance.RestaurantID = menurepo.IDRow(id)
	return b
}

func (b *MenuItemRowBuilder) Name(name string) *MenuItemRowBuilder {
	b.instance.Name = name
	return b
}

func (b *MenuItemRowBuilder) Price(price uint64) *MenuItemRowBuilder {
	b.instance.Price = menurepo.MoneyRow(price)
	return b
}

func (b *MenuItemRowBuilder) CreatedAt(createdAt time.Time) *MenuItemRowBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

func (b *MenuItemRowBuilder) UpdatedAt(updatedAt time.Time) *MenuItemRowBuilder {
	b.instance.UpdatedAt = sql.NullTime{Time: updatedAt, Valid: true}
	return b
}

func (b *MenuItemRowBuilder) DeletedAt(deletedAt time.Time) *MenuItemRowBuilder {
	b.instance.DeletedAt = sql.NullTime{Time: deletedAt, Valid: true}
	return b
}

func (b *MenuItemRowBuilder) build() *menurepo.MenuItemRow {
	if b.instance.CreatedAt.IsZero() {
		b.instance.CreatedAt = time.Now()
	}
	return b.instance
}

func (b *MenuItemRowBuilder) BuildP() *menurepo.MenuItemRow {
	return b.build()
}

func (b *MenuItemRowBuilder) BuildV() menurepo.MenuItemRow {
	return *(b.build())
}
