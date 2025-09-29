package graph

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/graphql-go/graphql"
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

func TestConvertArgsToFieldsets(t *testing.T) {
	t.Run("convertArgsToFieldsets should return empty list when args is null", func(tt *testing.T) {
		// arrange

		// act
		fieldsets, err := convertArgsToFieldsets(nil)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err == nil)
	})

	t.Run("convertArgsToFieldsets should return error when no array of type any", func(tt *testing.T) {
		// arrange
		arg := struct{}{}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err != nil)
	})

	t.Run("convertArgsToFieldsets should return error when fielset is not map", func(tt *testing.T) {
		// arrange
		arg := []any{
			struct{}{},
		}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err != nil)
		assert.Assert(tt, strings.Contains(err.Error(), "Error trying to convert fieldset"))
		assert.Assert(tt, strings.Contains(err.Error(), "to be a map[string]interface{}"))
	})

	t.Run("convertArgsToFieldsets should return error when inputs is not of type array any", func(tt *testing.T) {
		// arrange
		arg := []any{
			map[string]any{
				"inputs": struct{}{},
			},
		}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err != nil)
		assert.Assert(tt, strings.Contains(err.Error(), "Error trying to convert input"))
		assert.Assert(tt, strings.Contains(err.Error(), "to be an array type of any"))
	})

	t.Run("convertArgsToFieldsets should return error when cannot get input field as map[string]interface{}", func(tt *testing.T) {
		// arrange
		arg := []any{
			map[string]any{
				"inputs": []any{
					struct{}{},
				},
			},
		}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err != nil)
		assert.Assert(tt, strings.Contains(err.Error(), "Error trying to convert input"))
		assert.Assert(tt, strings.Contains(err.Error(), "to be a map[string]interface{}"))
	})

	t.Run("should return error when legend is not of type string", func(tt *testing.T) {
		// arrange
		arg := []any{
			map[string]any{
				"legend": struct{}{},
				"inputs": []any{
					map[string]any{
						"label": "",
					},
				},
			},
		}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 0)
		assert.Assert(tt, err != nil)
		assert.Assert(tt, strings.Contains(err.Error(), "Error trying to convert legend "))
		assert.Assert(tt, strings.Contains(err.Error(), "to be string"))
	})

	t.Run("should return nil error and fieldsets when data structure is corect", func(tt *testing.T) {
		// arrange
		arg := []any{
			map[string]any{
				"legend": "General Info",
				"inputs": []any{
					map[string]any{
						"label": "",
					},
				},
			},
		}

		// act
		fieldsets, err := convertArgsToFieldsets(arg)

		// assert
		assert.Assert(tt, len(fieldsets) == 1)
		assert.Assert(tt, err == nil)
	})
}

func TestCreateFormResolver(t *testing.T) {
	t.Run("createFormResolver should return error when repsitory is nil", func(tt *testing.T) {
		// arrange
		resolveParams := graphql.ResolveParams{
			Source:  nil,
			Args:    map[string]any{},
			Context: context.TODO(),
		}

		// act
		res, err := createFormResolver(resolveParams)

		// assert
		assert.Assert(tt, err != nil, "Error expected to be nil but got %s", err)
		assert.Assert(tt, res == nil, "Expect response to be nil but got %v", res)
	})

	t.Run("createFormResolver should return error when repsitory is nil", func(tt *testing.T) {
		// arrange
		resolveParams := graphql.ResolveParams{
			Source:  nil,
			Args:    map[string]any{},
			Context: context.TODO(),
		}

		// act
		res, err := createFormResolver(resolveParams)

		// assert
		assert.Assert(tt, err != nil, "Error expected to be nil but got %s", err)
		assert.Assert(tt, res == nil, "Expect response to be nil but got %v", res)
	})
}
