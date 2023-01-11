package storage_test

import (
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
	"github.com/pokrovsky-io/neows-asteroids/internal/repo/storage"
	"github.com/pokrovsky-io/neows-asteroids/pkg/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	dbUrl = "postgres://admin:admin@localhost:5432/asteroids?sslmode=disable"
)

var testCases = []entity.AsteroidsReport{
	{
		Date:  "2020-01-20",
		Count: 13,
	},
	{
		Date:  "2020-03-20",
		Count: 14,
	},
	{
		Date:  "2020-05-24",
		Count: 14,
	},
	{
		Date:  "2020-08-22",
		Count: 21,
	},
	{
		Date:  "2020-08-29",
		Count: 14,
	},
	{
		Date:  "2020-10-20",
		Count: 35,
	},
	{
		Date:  "2020-10-29",
		Count: 25,
	},
}

// TODO Вынести общие данные в подготовку

func TestStorage_CreateAndGet(t *testing.T) {
	db, err := postgres.New(dbUrl)
	if err != nil {
		t.Error(err)
	}

	stg, err := storage.New(db)
	if err != nil {
		t.Error(err)
	}

	if err = stg.Create(testCases); err != nil {
		t.Error(err)
	}

	var dates []string
	for _, tc := range testCases {
		dates = append(dates, tc.Date)
	}

	res, err := stg.Get(dates)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, testCases, res)
}

func TestStorage_Clear(t *testing.T) {
	db, err := postgres.New(dbUrl)
	if err != nil {
		t.Error(err)
	}

	stg, err := storage.New(db)
	if err != nil {
		t.Error(err)
	}

	if stg.GetLen() == 0 {
		t.Error("no elements")
	}

	if err = stg.Clear(); err != nil {
		t.Error(err)
	}

	assert.Equal(t, 0, stg.GetLen())
}
