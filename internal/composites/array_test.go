package composites_test

import (
	"testing"

	core "github.com/abyanmajid/v/internal"
	"github.com/abyanmajid/v/internal/composites"
	"github.com/abyanmajid/v/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewArraySchema(t *testing.T) {
	innerSchema := &core.Schema[int]{}
	arraySchema := composites.NewArraySchema("bruh", innerSchema)

	assert.NotNil(t, arraySchema)
	assert.Equal(t, "bruh", arraySchema.Schema.Path)
	assert.Equal(t, innerSchema, arraySchema.Inner)
}

func TestArraySchema_Parse(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema)

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

func TestArraySchema_ParseTyped(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema)

	t.Run("Valid typed array", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1, 2, 3})
		assert.True(t, result.Success)
		assert.Equal(t, []int{1, 2, 3}, result.Value)
	})
}

func TestArraySchema_Nonempty(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema).Nonempty()

	t.Run("Non-empty array", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1})
		assert.True(t, result.Success)
	})

	t.Run("Empty array", func(t *testing.T) {
		result := arraySchema.Parse([]int{})
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Array must not be empty")
	})
}

func TestArraySchema_Min(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema).Min(2)

	t.Run("Array meets min length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1, 2})
		assert.True(t, result.Success)
	})

	t.Run("Array does not meet min length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1})
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Array must have at least 2 elements")
	})
}

func TestArraySchema_Max(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema).Max(2)

	t.Run("Array meets max length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1, 2})
		assert.True(t, result.Success)
	})

	t.Run("Array exceeds max length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1, 2, 3})
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Array must have at most 2 elements")
	})
}

func TestArraySchema_Length(t *testing.T) {
	innerSchema := primitives.NewNumberSchema[int]("hi").Schema
	arraySchema := composites.NewArraySchema("bruh", innerSchema).Length(2)

	t.Run("Array meets exact length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1, 2})
		assert.True(t, result.Success)
	})

	t.Run("Array does not meet exact length", func(t *testing.T) {
		result := arraySchema.ParseTyped([]int{1})
		assert.False(t, result.Success)
		assert.Contains(t, result.Errors, "Array must have exactly 2 elements")
	})
}
