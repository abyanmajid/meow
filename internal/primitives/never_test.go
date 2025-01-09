package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNeverSchema(t *testing.T) {
	path := "hi"
	schema := NewNeverSchema(path)

	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Schema.Path)
	assert.Empty(t, schema.Schema.Rules)
}

func TestNeverSchema_Parse(t *testing.T) {
	schema := NewNeverSchema("hi")
	result := schema.Parse("any value")

	assert.NotNil(t, result)
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Value is not allowed.")
}
