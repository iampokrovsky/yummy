package fix

import (
	"database/sql"
	"time"
	menu_model "yummy/internal/app/menu/model"
)

type MenuItemRowBuilder struct {
	instance *menu_model.MenuItem
}

func MenuItemRow() *MenuItemRowBuilder {
	return &MenuItemRowBuilder{
		instance: &menu_model.MenuItem{},
	}
}

func (b *MenuItemRowBuilder) ID(id uint64) *MenuItemRowBuilder {
	b.instance.ID = menu_model.ID(id)
	return b
}

func (b *MenuItemRowBuilder) RestaurantID(id uint64) *MenuItemRowBuilder {
	b.instance.RestaurantID = menu_model.ID(id)
	return b
}

func (b *MenuItemRowBuilder) Name(name string) *MenuItemRowBuilder {
	b.instance.Name = name
	return b
}

func (b *MenuItemRowBuilder) Price(price uint64) *MenuItemRowBuilder {
	b.instance.Price = menu_model.Money(price)
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

func (b *MenuItemRowBuilder) build() *menu_model.MenuItem {
	if b.instance.CreatedAt.IsZero() {
		b.instance.CreatedAt = time.Now()
	}
	return b.instance
}

func (b *MenuItemRowBuilder) BuildP() *menu_model.MenuItem {
	return b.build()
}

func (b *MenuItemRowBuilder) BuildV() menu_model.MenuItem {
	return *(b.build())
}
