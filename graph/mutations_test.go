package graph

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestMapFieldSetInput(t *testing.T) {
	// arrange
	minLength := int(10)
	maxLength := int(43)
	args := map[string]any{
		"label":     nil,
		"type":      "datetime",
		"minLength": &minLength,
		"maxLength": &maxLength,
		"min":       "10",
		"max":       "10",
		"required":  true,
		"list":      nil,
		"multiple":  true,
	}

	// act
	model, err := mapFieldSetInput(args)

	// assert
	assert.NilError(t, err)
	assert.Assert(t, model.Label == "")
	assert.Assert(t, model.Type == "datetime")
	assert.Assert(t, model.MinLength == &minLength)
	assert.Assert(t, model.MaxLength == &maxLength)
	assert.Assert(t, model.Min == "10")
	assert.Assert(t, model.Max == "10")
	assert.Assert(t, model.Required == true)
	assert.Assert(t, model.Multiple == true)
}
