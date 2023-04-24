//go:build integration

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	menu_model "yummy/internal/app/menu/model"
	menu_repo "yummy/internal/app/menu/repo"
	"yummy/test/fix"
)

func TestMenuRepo_Create(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name  string
		items []menu_model.MenuItem
		ids   []uint64
		err   error
	}{
		{
			name: "success",
			items: []menu_model.MenuItem{
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
				fix.MenuItemRow().BuildV(),
			},
			ids: []uint64{1, 2, 3},
		},
		{
			name: "no input",
			err:  menu_repo.ErrBuildQuery,
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

func TestMenuRepo_ListByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		db.SetUp(ctx, t)
		defer db.TearDown(ctx)

		items := []menu_model.MenuItem{
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

	items := []menu_model.MenuItem{
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
		err    error
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
			err:    menu_repo.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resItems, err := menuRepo.ListByRestaurantID(ctx, tc.restId)
			assert.Equal(t, tc.err, err)

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

	items := []menu_model.MenuItem{
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
		err      error
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
			err:      menu_repo.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resItems, err := menuRepo.ListByName(ctx, tc.itemName)
			assert.Equal(t, tc.err, err)

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

	items := []menu_model.MenuItem{
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
		items []menu_model.MenuItem
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
			items: []menu_model.MenuItem{
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

	items := []menu_model.MenuItem{
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

	items := []menu_model.MenuItem{
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
