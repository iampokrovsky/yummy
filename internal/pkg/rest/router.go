package rest

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	menu_model "yummy/internal/app/menu/model"
)

// MenuRepo is a repository for menu items.
type MenuRepo interface {
	Create(ctx context.Context, items ...menu_model.MenuItem) ([]uint64, error)
	GetByID(ctx context.Context, id uint64) (menu_model.MenuItem, error)
	ListByRestaurantID(ctx context.Context, restId uint64) ([]menu_model.MenuItem, error)
	ListByName(ctx context.Context, name string) ([]menu_model.MenuItem, error)
	Update(ctx context.Context, items ...menu_model.MenuItem) (bool, error)
	Delete(ctx context.Context, ids ...uint64) (bool, error)
	Restore(ctx context.Context, ids ...uint64) (bool, error)
}

// Router is a router for repository methods of all services.
type Router struct {
	router   *httprouter.Router
	menuRepo MenuRepo
}

// NewRouter creates a new Router.
func NewRouter(menuRepo MenuRepo) *Router {
	r := Router{
		router:   httprouter.New(),
		menuRepo: menuRepo,
	}

	r.configureRoutes()

	return &r
}

func (rt *Router) configureRoutes() {
	rt.router.POST("/menu", rt.createMenuItems)
	rt.router.GET("/menu/item/:id", rt.getMenuItemByID)
	rt.router.GET("/menu/restaurant/:id", rt.listMenuItemsByRestaurantID)
	rt.router.GET("/menu/name/:name", rt.listMenuItemsByName)
	rt.router.PUT("/menu", rt.updateMenuItems)
	rt.router.DELETE("/menu/item/:id", rt.deleteMenuItem)
	rt.router.PATCH("/menu/item/:id/restore", rt.restoreMenuItem)
}

// ServeHTTP implements the http.Handler interface.
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.router.ServeHTTP(w, r)
}
