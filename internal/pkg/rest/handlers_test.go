package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	menu_model "yummy/internal/app/menu/model"
	menu_repo "yummy/internal/app/menu/repo"
	"yummy/test/fix"
	"yummy/test/mocks"
)

type menuRepoFix struct {
	ctrl   *gomock.Controller
	repo   *mocks.MockMenuRepo
	router *Router
}

func setUp(t *testing.T) menuRepoFix {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockMenuRepo(ctrl)
	router := NewRouter(repo)

	return menuRepoFix{
		ctrl:   ctrl,
		repo:   repo,
		router: router,
	}
}

func (fx *menuRepoFix) tearDown() {
	fx.ctrl.Finish()
}

func compareMenuItemRows(t *testing.T, exp, act []menu_model.MenuItem) (bool, error) {
	t.Helper()

	if len(exp) != len(act) {
		return false, errors.New("length of expected and actual results are not equal")
	}

	for i := range exp {
		if exp[i].ID != act[i].ID {
			return false, errors.New("id is not equal")
		}
		if exp[i].RestaurantID != act[i].RestaurantID {
			return false, errors.New("restaurant id is not equal")
		}
		if exp[i].Name != act[i].Name {
			return false, errors.New("name is not equal")
		}
		if exp[i].Price != act[i].Price {
			return false, errors.New("price is not equal")
		}
	}

	return true, nil
}

func TestRouter_createMenuItems(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	items := []interface{}{
		fix.MenuItemRow().RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	testCases := []struct {
		name   string
		isRepo bool
		ids    []uint64
		err    error
		data   []byte
		code   int
	}{
		{
			name:   "ok",
			isRepo: true,
			data: func() []byte {
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			ids:  []uint64{1, 2, 3},
			code: http.StatusCreated,
		},
		{
			name: "invalid data",
			data: []byte("bad json"),
			code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				fx.repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(tc.ids, tc.err)
			}

			req := httptest.NewRequest(http.MethodPost, "/menu", bytes.NewReader(tc.data))
			res := httptest.NewRecorder()

			// Act
			fx.router.createMenuItems(res, req, nil)

			// Assert
			// Check response status code
			assert.Equal(t, tc.code, res.Code)

			// Check ids
			if tc.isRepo {
				var resp struct {
					Ids []uint64 `json:"ids"`
				}

				if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, tc.ids, resp.Ids)
			}
		})
	}
}

func TestRouter_getMenuItemByID(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		isRepo bool
		id     string
		item   menu_model.MenuItem
		err    error
		code   int
		isBody bool
	}{
		{
			name:   "ok",
			isRepo: true,
			id:     "1",
			item:   fix.MenuItemRow().ID(1).RestaurantID(1).Name("test").Price(100000).BuildV(),
			code:   http.StatusOK,
			isBody: true,
		},
		{
			name: "invalid id",
			id:   "asd",
			code: http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusBadRequest,
		},
		{
			name:   "not found",
			id:     "123",
			err:    menu_repo.ErrNotFound,
			code:   http.StatusNotFound,
			isRepo: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				uint64Id, err := strconv.ParseUint(tc.id, 10, 64)
				if err != nil {
					t.Fatal(err)
				}
				fx.repo.EXPECT().GetByID(gomock.Any(), uint64Id).Return(tc.item, tc.err)
			}

			req := httptest.NewRequest(http.MethodGet, "/menu/item", nil)
			res := httptest.NewRecorder()
			params := []httprouter.Param{{Key: "id", Value: tc.id}}

			// Act
			fx.router.getMenuItemByID(res, req, params)

			// Assert

			// Check response status code
			assert.Equal(t, tc.code, res.Code)

			if tc.isBody {
				// Check response body
				var resItem menu_model.MenuItem
				if err := json.NewDecoder(res.Body).Decode(&resItem); err != nil {
					t.Fatal(err)
				}
				match, err := compareMenuItemRows(t, []menu_model.MenuItem{tc.item}, []menu_model.MenuItem{resItem})
				if !match && err != nil {
					t.Fatal(err)
				}
			}

		})
	}
}

func TestRouter_listMenuItemsByRestaurantID(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		isRepo bool
		restId string
		items  []menu_model.MenuItem
		err    error
		code   int
		isBody bool
	}{
		{
			name:   "ok",
			isRepo: true,
			restId: "1",
			items: []menu_model.MenuItem{
				fix.MenuItemRow().RestaurantID(1).Name("test1").Price(100000).BuildV(),
				fix.MenuItemRow().RestaurantID(1).Name("test2").Price(200000).BuildV(),
				fix.MenuItemRow().RestaurantID(1).Name("test3").Price(300000).BuildV(),
			},
			code:   http.StatusOK,
			isBody: true,
		},
		{
			name:   "invalid id",
			restId: "asd",
			code:   http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusBadRequest,
		},
		{
			name:   "not found",
			restId: "123",
			err:    menu_repo.ErrNotFound,
			code:   http.StatusNotFound,
			isRepo: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				uint64Id, err := strconv.ParseUint(tc.restId, 10, 64)
				if err != nil {
					t.Fatal(err)
				}
				fx.repo.EXPECT().ListByRestaurantID(gomock.Any(), uint64Id).Return(tc.items, tc.err)
			}

			req := httptest.NewRequest(http.MethodGet, "/menu/restaurant", nil)
			res := httptest.NewRecorder()
			params := []httprouter.Param{{Key: "id", Value: tc.restId}}

			// Act
			fx.router.listMenuItemsByRestaurantID(res, req, params)

			// Assert

			// Check response status code
			assert.Equal(t, tc.code, res.Code)

			if tc.isBody {
				// Check response body
				var resItems []menu_model.MenuItem
				if err := json.NewDecoder(res.Body).Decode(&resItems); err != nil {
					t.Fatal(err)
				}
				match, err := compareMenuItemRows(t, tc.items, resItems)
				if !match && err != nil {
					t.Fatal(err)
				}
			}

		})
	}
}

func TestRouter_listMenuItemsByName(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name     string
		isRepo   bool
		itemName string
		items    []menu_model.MenuItem
		err      error
		code     int
		isBody   bool
	}{
		{
			name:     "ok",
			isRepo:   true,
			itemName: "test",
			items: []menu_model.MenuItem{
				fix.MenuItemRow().RestaurantID(1).Name("test1").Price(100000).BuildV(),
				fix.MenuItemRow().RestaurantID(1).Name("test2").Price(200000).BuildV(),
				fix.MenuItemRow().RestaurantID(1).Name("test3").Price(300000).BuildV(),
			},
			code:   http.StatusOK,
			isBody: true,
		},
		{
			name: "empty",
			code: http.StatusBadRequest,
		},
		{
			name:     "not found",
			itemName: "asd",
			err:      menu_repo.ErrNotFound,
			code:     http.StatusNotFound,
			isRepo:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				fx.repo.EXPECT().ListByName(gomock.Any(), tc.itemName).Return(tc.items, tc.err)
			}

			req := httptest.NewRequest(http.MethodGet, "/menu/name", nil)
			res := httptest.NewRecorder()
			params := []httprouter.Param{{Key: "name", Value: tc.itemName}}

			// Act
			fx.router.listMenuItemsByName(res, req, params)

			// Assert

			// Check response status code
			assert.Equal(t, tc.code, res.Code)

			if tc.isBody {
				// Check response body
				var resItems []menu_model.MenuItem
				if err := json.NewDecoder(res.Body).Decode(&resItems); err != nil {
					t.Fatal(err)
				}
				match, err := compareMenuItemRows(t, tc.items, resItems)
				if !match && err != nil {
					t.Fatal(err)
				}
			}

		})
	}
}

func TestRouter_updateMenuItems(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		isRepo bool
		data   []byte
		repoOk bool
		code   int
	}{
		{
			name:   "ok",
			isRepo: true,
			data: func() []byte {
				items := []interface{}{
					fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
					fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
					fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
				}
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			repoOk: true,
			code:   http.StatusOK,
		},
		{
			name:   "objects not found",
			isRepo: true,
			data: func() []byte {
				items := []interface{}{
					fix.MenuItemRow().RestaurantID(1).Name("test4").Price(400000).BuildV(),
					fix.MenuItemRow().RestaurantID(1).Name("test5").Price(500000).BuildV(),
					fix.MenuItemRow().RestaurantID(1).Name("test6").Price(600000).BuildV(),
				}
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			code: http.StatusNotFound,
		},
		{
			name: "invalid data",
			data: []byte("bad json"),
			code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				fx.repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(tc.repoOk, nil)
			}

			req := httptest.NewRequest(http.MethodPut, "/menu", bytes.NewReader(tc.data))
			res := httptest.NewRecorder()

			// Act
			fx.router.updateMenuItems(res, req, nil)

			// Assert
			// Check response status code
			assert.Equal(t, tc.code, res.Code)
		})
	}
}

func TestRouter_deleteMenuItem(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		isRepo bool
		id     string
		repoOk bool
		code   int
	}{
		{
			name:   "ok",
			isRepo: true,
			id:     "1",
			repoOk: true,
			code:   http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "asd",
			code: http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusBadRequest,
		},
		{
			name:   "not found",
			isRepo: true,
			id:     "123",
			repoOk: false,
			code:   http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				uint64Id, err := strconv.ParseUint(tc.id, 10, 64)
				if err != nil {
					t.Fatal(err)
				}
				fx.repo.EXPECT().Delete(gomock.Any(), uint64Id).Return(tc.repoOk, nil)
			}

			req := httptest.NewRequest(http.MethodDelete, "/menu/item", nil)
			res := httptest.NewRecorder()
			params := []httprouter.Param{{Key: "id", Value: tc.id}}

			// Act
			fx.router.deleteMenuItem(res, req, params)

			// Assert
			// Check response status code
			assert.Equal(t, tc.code, res.Code)
		})
	}
}

func TestRouter_restoreMenuItem(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		isRepo bool
		id     string
		repoOk bool
		code   int
	}{
		{
			name:   "ok",
			isRepo: true,
			id:     "1",
			repoOk: true,
			code:   http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "asd",
			code: http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusBadRequest,
		},
		{
			name:   "not found",
			isRepo: true,
			id:     "123",
			repoOk: false,
			code:   http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			if tc.isRepo {
				uint64Id, err := strconv.ParseUint(tc.id, 10, 64)
				if err != nil {
					t.Fatal(err)
				}
				fx.repo.EXPECT().Restore(gomock.Any(), uint64Id).Return(tc.repoOk, nil)
			}

			req := httptest.NewRequest(http.MethodDelete, "/menu/item/"+tc.id+"/restore", nil)
			res := httptest.NewRecorder()
			params := []httprouter.Param{{Key: "id", Value: tc.id}}

			// Act
			fx.router.restoreMenuItem(res, req, params)

			// Assert
			// Check response status code
			assert.Equal(t, tc.code, res.Code)
		})
	}
}
