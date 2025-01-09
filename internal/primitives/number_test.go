package primitives_test

import (
	"math"
	"testing"

	"github.com/abyanmajid/z/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewNumberSchema(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad")
	assert.NotNil(t, schema)
	assert.Equal(t, "im sad", schema.Base.Path)
	assert.Empty(t, schema.Base.Rules)
}

func TestNumberSchema_Parse(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad")

	result := schema.Parse(123.45)
	assert.True(t, result.Success)

	result = schema.Parse("not a number")
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a string.", result.Errors[0])
}

func TestNumberSchema_ParseTyped(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad")

	result := schema.ParseTyped(123.45)
	assert.True(t, result.Success)
}

func TestNumberSchema_Gt(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Gt(10)

	result := schema.ParseTyped(15)
	assert.True(t, result.Success)

	result = schema.ParseTyped(5)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be greater than 10", result.Errors[0])
}

func TestNumberSchema_Gte(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Gte(10)

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(5)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be greater than or equal to 10", result.Errors[0])
}

func TestNumberSchema_Lt(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Lt(10)

	result := schema.ParseTyped(5)
	assert.True(t, result.Success)

	result = schema.ParseTyped(15)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be smaller than 10", result.Errors[0])
}

func TestNumberSchema_Lte(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Lte(10)

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(15)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be smaller than or equal to 10", result.Errors[0])
}

func TestNumberSchema_Int(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Int()

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(10.5)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be an integer", result.Errors[0])
}

func TestNumberSchema_Positive(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Positive()

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(-10)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a positive number", result.Errors[0])
}

func TestNumberSchema_NonNegative(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").NonNegative()

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(-10)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a non-negative number", result.Errors[0])
}

func TestNumberSchema_Negative(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Negative()

	result := schema.ParseTyped(-10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(10)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a negative number", result.Errors[0])
}

func TestNumberSchema_NonPositive(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").NonPositive()

	result := schema.ParseTyped(-10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(10)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a non-positive number", result.Errors[0])
}

func TestNumberSchema_MultipleOf(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").MultipleOf(5)

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(7)
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a multiple of 5", result.Errors[0])
}

func TestNumberSchema_Finite(t *testing.T) {
	schema := primitives.NewNumberSchema("im sad").Finite()

	result := schema.ParseTyped(10)
	assert.True(t, result.Success)

	result = schema.ParseTyped(math.Inf(1))
	assert.False(t, result.Success)
	assert.Equal(t, "Must be a finite number", result.Errors[0])
}
