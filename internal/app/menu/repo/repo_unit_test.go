package repo

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock_postgres "yummy/pkg/postgres/mock"
)

type menuRepoFixture struct {
	ctrl   *gomock.Controller
	mockDB *mock_postgres.MockDB
	repo   *MenuRepo
}

func setUp(t *testing.T) menuRepoFixture {
	ctrl := gomock.NewController(t)
	mockDB := mock_postgres.NewMockDB(ctrl)
	repo := NewMenuRepo(mockDB)
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	mockDB.EXPECT().Builder().Return(&builder).AnyTimes()

	return menuRepoFixture{
		ctrl:   ctrl,
		mockDB: mockDB,
		repo:   repo,
	}
}

func (u *menuRepoFixture) tearDown() {
	u.ctrl.Finish()
}

func TestMenuRepo_Create(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name      string
		buildFail bool
		dbErr     error
		items     []MenuItemRow
		ids       []uint64
		expErr    error
	}{
		{
			name:  "success",
			items: []MenuItemRow{{}},
			ids:   []uint64{},
		},
		{
			name:      "failed to build query",
			buildFail: true,
			expErr:    ErrBuildQuery,
		},
		{
			name:   "internal error",
			dbErr:  assert.AnError,
			items:  []MenuItemRow{{}},
			expErr: assert.AnError,
		},
		{
			name:   "not found",
			dbErr:  sql.ErrNoRows,
			items:  []MenuItemRow{{}},
			expErr: ErrObjectNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qb := fx.mockDB.Builder().
				Insert("menu_items").
				Columns("restaurant_id", "name", "price")
			for _, item := range tc.items {
				qb = qb.Values(item.RestaurantID, item.Name, item.Price)
			}
			qb.Suffix(`RETURNING id`)
			query, _, _ := qb.ToSql()

			if !tc.buildFail {
				fx.mockDB.EXPECT().
					Query(gomock.Any(), query, gomock.Any()).
					Return(pgxpoolmock.NewRows([]string{}).ToPgxRows(), tc.dbErr)
			}

			ids, err := fx.repo.Create(context.Background(), tc.items)

			assert.Equal(t, tc.ids, ids)
			assert.Equal(t, tc.expErr, err)
		})
	}
}

func TestMenuRepo_GetByID(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name   string
		dbErr  error
		expErr error
	}{
		{
			name: "success",
		},
		{
			name:   "internal error",
			dbErr:  assert.AnError,
			expErr: assert.AnError,
		},
		{
			name:   "not found",
			dbErr:  sql.ErrNoRows,
			expErr: ErrObjectNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, _, _ := fx.mockDB.Builder().
				Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
				From("menu_items").
				Where(squirrel.Eq{"id": 0}).ToSql()

			fx.mockDB.EXPECT().Get(gomock.Any(), gomock.Any(), query, gomock.Any()).Return(tc.dbErr)

			item, err := fx.repo.GetByID(context.Background(), 0)

			assert.Equal(t, MenuItemRow{}, item)
			assert.Equal(t, tc.expErr, err)
		})
	}
}

func TestMenuRepo_ListByRestaurantID(t *testing.T) {
	t.Parallel()

	fx := setUp(t)
	defer fx.tearDown()

	testCases := []struct {
		name      string
		buildFail bool
		dbErr     error
		items     []MenuItemRow
		expErr    error
	}{
		{
			name: "success",
		},
		{
			name:   "internal error",
			dbErr:  assert.AnError,
			expErr: assert.AnError,
		},
		{
			name:   "not found",
			dbErr:  sql.ErrNoRows,
			expErr: ErrObjectNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, _, _ := fx.mockDB.Builder().
				Select("id", "restaurant_id", "name", "price", "created_at", "updated_at", "deleted_at").
				From("menu_items").
				Where(squirrel.Eq{"restaurant_id": 0}).ToSql()

			fx.mockDB.EXPECT().Select(gomock.Any(), gomock.Any(), query, gomock.Any()).Return(tc.dbErr)

			items, err := fx.repo.ListByRestaurantID(context.Background(), 0)

			assert.Equal(t, tc.items, items)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
