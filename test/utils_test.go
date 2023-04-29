package test

import (
	"errors"
	"testing"
	menu_model "yummy/internal/app/menu/model"
)

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
