package composites_test

import (
	"testing"

	core "github.com/abyanmajid/z/internal"
	"github.com/abyanmajid/z/internal/composites"
	"github.com/abyanmajid/z/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewArraySchema(t *testing.T) {
	innerSchema := &core.Schema[int]{}
	arraySchema := composites.NewArraySchema("testPath", innerSchema)

	assert.NotNil(t, arraySchema)
	assert.Equal(t, "testPath", arraySchema.Base.Path)
	assert.Equal(t, innerSchema, arraySchema.Inner)
}

func TestArraySchema_Parse(t *testing.T) {
	innerSchema := primitives.NewNumberSchema("hi").Int().Base
	arraySchema := composites.NewArraySchema("testPath", innerSchema)

	t.Run("Valid array", func(t *testing.T) {
		result := arraySchema.Parse([]interface{}{1, 2, 3})
		assert.True(t, result.Success)
		assert.Equal(t, []int{1, 2, 3}, result.Value)
	})

	t.Run("Invalid array element type", func(t *testing.T) {
		result := arraySchema.Parse([]interface{}{1, "two", 3})
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Element at index 1 must be of type int")
	})

	t.Run("Not an array", func(t *testing.T) {
		result := arraySchema.Parse("not an array")
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Must be an array")
	})
}
