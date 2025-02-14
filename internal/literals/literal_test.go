package literals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLiteralSchema(t *testing.T) {
	schema := NewLiteralSchema("test/path", "testValue")

	assert.NotNil(t, schema)
	assert.Equal(t, "test/path", schema.Schema.Path)
	assert.Equal(t, "testValue", schema.Value)
	assert.Empty(t, schema.Schema.Rules)
}

func TestLiteralSchema_Parse(t *testing.T) {
	schema := NewLiteralSchema("test/path", "testValue")

	t.Run("Valid value", func(t *testing.T) {
		result := schema.Parse("testValue")
		assert.True(t, result.Ok)
		assert.Empty(t, result.Errors)
	})

	t.Run("Invalid type", func(t *testing.T) {
		result := schema.Parse(123)
		assert.False(t, result.Ok)
		assert.Equal(t, "Invalid type.", result.Errors[0])
	})

	t.Run("Invalid value", func(t *testing.T) {
		result := schema.Parse("wrongValue")
		assert.False(t, result.Ok)
		assert.Equal(t, "Value must be testValue.", result.Errors[0])
	})
}
