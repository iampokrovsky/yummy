package model

import (
	"database/sql"
	"time"
)

type Cuisine string

const (
	CuisineItalian  Cuisine = "Italian"
	CuisineChinese          = "Chinese"
	CuisineMexican          = "Mexican"
	CuisineJapanese         = "Japanese"
	CuisineIndian           = "Indian"
	CuisineThai             = "Thai"
	CuisineFrench           = "French"
	CuisineGreek            = "Greek"
	CuisineKorean           = "Korean"
	CuisineRussian          = "Russian"
	CuisineGeorgian         = "Georgian"
)

type ID int64

type Restaurant struct {
	ID        ID           `db:"id"`
	Name      string       `db:"name"`
	Address   string       `db:"address"`
	Cuisine   Cuisine      `db:"cuisine"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
