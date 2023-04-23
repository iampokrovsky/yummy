// TODO: add //go:build integration
package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	menurepo "yummy/internal/app/menu/repo"
	"yummy/test/fix"
	"yummy/test/postgres"
)

var (
	dsn      = "host=localhost port=5432 user=user password=pass dbname=yummy_db sslmode=disable"
	db       *postgres.PostgresTestDB
	menuRepo *menurepo.MenuRepo
)

func init() {
	db = postgres.NewPostgresTestDB(context.Background(), dsn)
	menuRepo = menurepo.NewMenuRepo(db)
}

func compareUint64Slices(t *testing.T, a, b []uint64) bool {
	t.Helper()

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func compareMenuItemRows(t *testing.T, exp, act []menurepo.MenuItemRow) (bool, error) {
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

func TestMenuRepo_Create(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name  string
		items []menurepo.MenuItemRow
		ids   []uint64
		err   error
	}{
		{
			name: "success",
			items: []menurepo.MenuItemRow{
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
			},
			ids: []uint64{1, 2, 3},
		},
		{
			name: "no input",
			err:  menurepo.ErrBuildQuery,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			// act
			res, err := menuRepo.Create(ctx, tc.items...)

			// assert
			assert.Equal(t, tc.ids, res)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestMenuRepo_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		db.SetUp(ctx, t)
		defer db.TearDown(ctx)

		item := fix.MenuItemRow().ID(1).RestaurantID(1).Name("test").Price(100000).BuildV()

		_, err := menuRepo.Create(ctx, item)
		if err != nil {
			t.Fatal(err)
		}

		resItem, err := menuRepo.GetByID(ctx, uint64(item.ID))
		assert.Equal(t, err, nil)

		match, err := compareMenuItemRows(t, []menurepo.MenuItemRow{item}, []menurepo.MenuItemRow{resItem})
		if !match && err != nil {
			t.Fatal(err)
		}
	})
}

func TestMenuRepo_ListByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		db.SetUp(ctx, t)
		defer db.TearDown(ctx)

		items := []menurepo.MenuItemRow{
			fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
			fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
			fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
		}

		ids, err := menuRepo.Create(ctx, items...)
		if err != nil {
			t.Fatal(err)
		}

		resItems, err := menuRepo.ListByID(ctx, ids...)
		assert.Equal(t, err, nil)

		match, err := compareMenuItemRows(t, items, resItems)
		if !match && err != nil {
			t.Fatal(err)
		}
	})
}

func TestMenuRepo_ListByRestaurantID(t *testing.T) {
	ctx := context.Background()

	db.SetUp(ctx, t)
	defer db.TearDown(ctx)

	restId := uint64(1)

	items := []menurepo.MenuItemRow{
		fix.MenuItemRow().ID(1).RestaurantID(restId).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(restId).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(restId).Name("test3").Price(300000).BuildV(),
	}

	_, err := menuRepo.Create(ctx, items...)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name   string
		restId uint64
		match  bool
	}{
		{
			name:   "success",
			restId: restId,
			match:  true,
		},
		{
			name:   "not found",
			restId: 2,
			match:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resItems, err := menuRepo.ListByRestaurantID(ctx, tc.restId)
			if err != nil {
				t.Fatal(err)
			}
			assert.Nil(t, err)

			match, err := compareMenuItemRows(t, items, resItems)
			if tc.match != match && err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMenuRepo_ListByName(t *testing.T) {
	ctx := context.Background()

	db.SetUp(ctx, t)
	defer db.TearDown(ctx)

	name := "test"

	items := []menurepo.MenuItemRow{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name(name + "1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name(name + "2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name(name + "3").Price(300000).BuildV(),
	}

	_, err := menuRepo.Create(ctx, items...)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name     string
		itemName string
		match    bool
	}{
		{
			name:     "success",
			itemName: name,
			match:    true,
		},
		{
			name:     "not found",
			itemName: "asd",
			match:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resItems, err := menuRepo.ListByName(ctx, tc.itemName)
			if err != nil {
				t.Fatal(err)
			}
			assert.Nil(t, err)

			match, err := compareMenuItemRows(t, items, resItems)
			if tc.match != match && err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMenuRepo_Update(t *testing.T) {
	ctx := context.Background()

	db.SetUp(ctx, t)
	defer db.TearDown(ctx)

	items := []menurepo.MenuItemRow{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).BuildV(),
	}

	_, err := menuRepo.Create(ctx, items...)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name  string
		items []menurepo.MenuItemRow
		match bool
		isErr bool
	}{
		{
			name:  "success",
			items: items,
			match: true,
		},
		{
			name:  "empty input",
			match: true,
			isErr: true,
		},
		{
			name: "empty items fields",
			items: []menurepo.MenuItemRow{
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
			},
			match: false,
			isErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := menuRepo.Update(ctx, tc.items...)
			if tc.isErr && err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.match, res)
		})
	}
}

func TestMenuRepo_Delete(t *testing.T) {

	ctx := context.Background()

	items := []menurepo.MenuItemRow{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).DeletedAt(time.Now()).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).DeletedAt(time.Now()).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).DeletedAt(time.Now()).BuildV(),
	}

	testCases := []struct {
		name string
		ids  []uint64
		ok   bool
	}{
		{
			name: "success",
			ids:  []uint64{1, 2, 3},
			ok:   true,
		},
		{
			name: "no input",
			ok:   true,
		},
		{
			name: "bad ids",
			ids:  []uint64{4, 5, 6},
			ok:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			ids, err := menuRepo.Create(ctx, items...)
			if err != nil {
				t.Fatal(err)
			}

			ok, err := menuRepo.Delete(ctx, tc.ids...)
			assert.Equal(t, tc.ok, ok)
			assert.Nil(t, err)

			if compareUint64Slices(t, ids, tc.ids) {
				resItems, err := menuRepo.ListByID(ctx, ids...)
				if err != nil {
					t.Fatal(err)
				}

				match, err := compareMenuItemRows(t, items, resItems)
				if !match && err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestMenuRepo_Restore(t *testing.T) {

	ctx := context.Background()

	items := []menurepo.MenuItemRow{
		fix.MenuItemRow().ID(1).RestaurantID(1).Name("test1").Price(100000).DeletedAt(time.Now()).BuildV(),
		fix.MenuItemRow().ID(2).RestaurantID(1).Name("test2").Price(200000).DeletedAt(time.Now()).BuildV(),
		fix.MenuItemRow().ID(3).RestaurantID(1).Name("test3").Price(300000).DeletedAt(time.Now()).BuildV(),
	}

	testCases := []struct {
		name string
		ids  []uint64
		ok   bool
	}{
		{
			name: "success",
			ids:  []uint64{1, 2, 3},
			ok:   true,
		},
		{
			name: "no input",
			ok:   true,
		},
		{
			name: "bad ids",
			ids:  []uint64{4, 5, 6},
			ok:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db.SetUp(ctx, t)
			defer db.TearDown(ctx)

			ids, err := menuRepo.Create(ctx, items...)
			if err != nil {
				t.Fatal(err)
			}

			ok, err := menuRepo.Restore(ctx, tc.ids...)
			assert.Equal(t, tc.ok, ok)
			assert.Nil(t, err)

			if compareUint64Slices(t, ids, tc.ids) {
				resItems, err := menuRepo.ListByID(ctx, ids...)
				if err != nil {
					t.Fatal(err)
				}

				match, err := compareMenuItemRows(t, items, resItems)
				if !match && err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
