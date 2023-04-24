//go:build integration

package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	menu_model "yummy/internal/app/menu/model"
	"yummy/test/fix"
)

func TestAPI_createMenuItems(t *testing.T) {
	ctx := context.Background()

	items := []interface{}{
		fix.MenuItemRow().RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	testCases := []struct {
		name   string
		ids    []uint64
		data   []byte
		code   int
		isBody bool
	}{
		{
			name: "ok",
			data: func() []byte {
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			ids:    []uint64{1, 2, 3},
			code:   http.StatusCreated,
			isBody: true,
		},
		{
			name: "invalid data",
			data: []byte("bad json"),
			code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Get items
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(tc.data))

			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			assert.Equal(t, tc.code, res.StatusCode)

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)

			if tc.isBody {
				// Check ids
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

func TestAPI_getMenuItemByID(t *testing.T) {
	ctx := context.Background()

	item := fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV()

	testCases := []struct {
		name   string
		id     string
		data   []byte
		code   int
		isBody bool
		item   menu_model.MenuItem
	}{
		{
			name: "ok",
			id:   "1",
			data: func() []byte {
				b, err := json.Marshal([]menu_model.MenuItem{item})
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			code:   http.StatusOK,
			isBody: true,
			item:   item,
		},
		{
			name: "invalid id",
			id:   "asd",
			code: http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusNotFound,
		},
		{
			name: "not found",
			id:   "123",
			code: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Get items
			res, err = http.Get(server.URL + "/menu/item/" + tc.id)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)

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

func TestAPI_listMenuItemsByRestaurantID(t *testing.T) {
	ctx := context.Background()

	items := []menu_model.MenuItem{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	testCases := []struct {
		name   string
		items  []menu_model.MenuItem
		data   []byte
		restId string
		code   int
		isBody bool
	}{
		{
			name:   "ok",
			restId: "1",
			items:  items,
			data: func() []byte {
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
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
			code: http.StatusNotFound,
		},
		{
			name:   "not found",
			restId: "123",
			code:   http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Get items
			res, err = http.Get(server.URL + "/menu/restaurant/" + tc.restId)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)

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

func TestAPI_listMenuItemsByName(t *testing.T) {
	ctx := context.Background()

	items := []menu_model.MenuItem{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	testCases := []struct {
		name     string
		items    []menu_model.MenuItem
		data     []byte
		itemName string
		code     int
		isBody   bool
	}{
		{
			name:  "ok",
			items: items,
			data: func() []byte {
				b, err := json.Marshal(items)
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			itemName: "test",
			code:     http.StatusOK,
			isBody:   true,
		},
		{
			name: "empty request",
			code: http.StatusNotFound,
		},
		{
			name:     "not found",
			itemName: "asd",
			code:     http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Get items
			res, err = http.Get(server.URL + "/menu/name/" + tc.itemName)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)

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

func TestAPI_updateMenuItems(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name string
		data []byte
		code int
	}{
		{
			name: "ok",
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
			code: http.StatusOK,
		},
		{
			name: "objects not found",
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
			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Prepare request
			req, err := http.NewRequest(http.MethodPut, server.URL+"/menu", bytes.NewBuffer(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Send request
			client := &http.Client{}
			res, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)
		})
	}
}

func TestAPI_deleteMenuItem(t *testing.T) {
	ctx := context.Background()

	items := []menu_model.MenuItem{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	encodedItems, err := json.Marshal(items)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name string
		id   string
		code int
	}{
		{
			name: "ok",
			id:   "1",
			code: http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "asd",
			code: http.StatusBadRequest,
		},
		{
			name: "empty",
			code: http.StatusNotFound,
		},
		{
			name: "not found",
			id:   "123",
			code: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(encodedItems))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Prepare request
			req, err := http.NewRequest(http.MethodDelete, server.URL+"/menu/item/"+tc.id, bytes.NewBuffer(encodedItems))
			if err != nil {
				t.Fatal(err)
			}

			// Send request
			client := &http.Client{}
			res, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)
		})
	}
}

func TestAPI_restoreMenuItems(t *testing.T) {
	ctx := context.Background()

	items := []menu_model.MenuItem{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	encodedItems, err := json.Marshal(items)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name string
		id   string
		code int
	}{
		{
			name: "ok",
			id:   "1",
			code: http.StatusOK,
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
			name: "not found",
			id:   "123",
			code: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup test database
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// Insert data to database
			res, err := http.Post(server.URL+"/menu", "application/json", bytes.NewReader(encodedItems))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Prepare request
			req, err := http.NewRequest(http.MethodPatch, server.URL+"/menu/item/"+tc.id+"/restore", bytes.NewBuffer(encodedItems))
			if err != nil {
				t.Fatal(err)
			}

			// Send request
			client := &http.Client{}
			res, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// Check response status code
			assert.Equal(t, tc.code, res.StatusCode)
		})
	}
}
