package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNilSchema(t *testing.T) {
	path := "bruh"
	schema := NewNilSchema(path)

	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Base.Path)
	assert.Empty(t, schema.Base.Rules)
}

func TestNilSchema_Parse(t *testing.T) {
	path := "bruh"
	schema := NewNilSchema(path)

	t.Run("Value is nil", func(t *testing.T) {
		result := schema.Parse(nil)
		assert.True(t, result.Success)
		assert.Empty(t, result.Errors)
	})

	t.Run("Value is not nil", func(t *testing.T) {
		result := schema.Parse("not nil")
		assert.False(t, result.Success)
		assert.Equal(t, "Value must be nil.", result.Errors[0])
	})
}
