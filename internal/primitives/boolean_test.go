package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBooleanSchema(t *testing.T) {
	path := "abyan has a majestic cat"
	schema := NewBooleanSchema(path)

	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Schema.Path)

	assert.Empty(t, schema.Schema.Rules)
}

func TestBooleanSchema_Parse(t *testing.T) {
	schema := NewBooleanSchema("abyan has a majestic cat")

	t.Run("valid boolean", func(t *testing.T) {
		result := schema.Parse(true)
		assert.True(t, result.Ok)
		assert.Equal(t, true, result.Value)
	})

	t.Run("invalid boolean", func(t *testing.T) {
		result := schema.Parse("not a boolean")
		assert.False(t, result.Ok)
		assert.Contains(t, result.Errors, "Must be a boolean")
	})
}

func TestBooleanSchema_ParseTyped(t *testing.T) {
	schema := NewBooleanSchema("abyan has a majestic cat")

	result := schema.ParseTyped(true)
	assert.True(t, result.Ok)
	assert.Equal(t, true, result.Value)
}
