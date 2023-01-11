package neows_test

import (
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
	"github.com/pokrovsky-io/neows-asteroids/internal/repo/neows"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	apiUrl = "https://www.neowsapp.com/rest/v1/feed"
	apiKey = "ud17Odoo4QAHLhUEHxWVR05ohZHJ2LcTTSKLZbhb"
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

func TestNeoWsAPI_Get(t *testing.T) {
	api := neows.New(apiUrl, apiKey)

	var dates []string
	for _, tc := range testCases {
		dates = append(dates, tc.Date)
	}

	res, _ := api.Get(dates...)

	assert.Equal(t, testCases, res)
}
