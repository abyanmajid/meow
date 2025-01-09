package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnySchema(t *testing.T) {
	path := "lol"
	schema := NewAnySchema(path)

	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Schema.Path)
	assert.Empty(t, schema.Schema.Rules)
}

func TestAnySchema_Parse(t *testing.T) {
	schema := NewAnySchema("lol")
	value := "test value"

	result := schema.Parse(value)

	assert.NotNil(t, result)
	assert.Equal(t, value, result.Value)
	assert.True(t, result.Success)
}
