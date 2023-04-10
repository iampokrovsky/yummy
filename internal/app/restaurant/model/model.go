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
	ID        ID           `db:"id" json:"id"`
	Name      string       `db:"name" json:"name"`
	Address   string       `db:"address" json:"address"`
	Cuisine   Cuisine      `db:"cuisine" json:"cuisine"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updatedAt"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deletedAt"`
}
