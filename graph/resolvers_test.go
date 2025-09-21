package graph

import (
	"slices"
	"testing"

	"github.com/squishedfox/fictional-fiesta/db"
	"gotest.tools/v3/assert"
)

func TestGetFormsModel(t *testing.T) {
	t.Run("getFormsModel should use default limit when not provided in args", func(tt *testing.T) {
		// arrange
		args := map[string]any{}

		// act
		model := getFormsModel(args)

		// assert
		assert.Assert(tt, model.Limit == int64(DEFAULT_LIMIT_VALUE))
		assert.Assert(tt, len(model.Filters) == 0, "Expected filters to be 0 but got %d", len(model.Filters))
	})

	t.Run("getFormsModel should use default page when not provided in args", func(tt *testing.T) {
		// arrange
		args := map[string]any{}

		// act
		model := getFormsModel(args)

		// assert
		assert.Assert(tt, model.Skip == 0, "Expected Skip to be 0 but got %d", model.Skip)
		assert.Assert(tt, len(model.Filters) == 0, "Expected filters to be 0 but got %d", len(model.Filters))
	})

	t.Run("getFormsModel should set skip when limit and page provided", func(tt *testing.T) {
		// arrange
		args := map[string]any{
			"limit": 2,
			"page":  3, // indexing based on 0, 1, 2
		}

		// act
		model := getFormsModel(args)

		// assert
		assert.Assert(tt, model.Skip == 6, "Expected Skip to be 6 but got %d", model.Skip)
		assert.Assert(tt, len(model.Filters) == 0, "Expected filters to be 0 but got %d", len(model.Filters))
	})

	t.Run("getFormsModel should set name when provided in args", func(tt *testing.T) {
		// arrange
		args := map[string]any{
			"name": "A funny form",
		}

		// act
		model := getFormsModel(args)

		// assert
		assert.Assert(tt, model.Skip == 0, "Expected Skip to be 0 but got %d", model.Skip)
		assert.Assert(tt, len(model.Filters) == 1, "Expected filters to be 1 but got %d", len(model.Filters))
		containsfieldFilter := slices.ContainsFunc(model.Filters, func(filter *db.SearchFieldFilter) bool {
			return filter.Field == "name"
		})
		assert.Assert(tt, containsfieldFilter == true, "Expected filters to contain name field but did not")
	})
}
