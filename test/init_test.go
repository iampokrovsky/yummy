package test

import (
	"context"
	"net/http/httptest"
	menu_repo "yummy/internal/app/menu/repo"
	"yummy/internal/pkg/rest"
	"yummy/test/postgres"
)

var (
	dsn      = "host=localhost port=5432 user=user password=pass dbname=yummy_db sslmode=disable"
	db       *postgres.PostgresTestDB
	menuRepo *menu_repo.MenuRepository
	server   *httptest.Server
)

func init() {
	db = postgres.NewPostgresTestDB(context.Background(), dsn)
	menuRepo = menu_repo.NewMenuRepo(db)
	router := rest.NewRouter(menuRepo)
	server = httptest.NewServer(router)
}
