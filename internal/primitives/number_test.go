package primitives_test

import (
	"math"
	"testing"

	"github.com/abyanmajid/v/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewNumberSchema(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat")
	assert.NotNil(t, schema)
	assert.Equal(t, "abyan has a majestic cat", schema.Schema.Path)
	assert.Empty(t, schema.Schema.Rules)
}

func TestNumberSchema_Parse(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat")

	result := schema.Parse(123.45)
	assert.True(t, result.Ok)

	result = schema.Parse("not a number")
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a number.", result.Errors[0])
}

func TestNumberSchema_ParseTyped(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat")

	result := schema.ParseTyped(123.45)
	assert.True(t, result.Ok)
}

func TestNumberSchema_Gt(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Gt(10)

	result := schema.ParseTyped(15)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(5)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be greater than 10", result.Errors[0])
}

func TestNumberSchema_Gte(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Gte(10)

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(5)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be greater than or equal to 10", result.Errors[0])
}

func TestNumberSchema_Lt(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Lt(10)

	result := schema.ParseTyped(5)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(15)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be smaller than 10", result.Errors[0])
}

func TestNumberSchema_Lte(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Lte(10)

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(15)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be smaller than or equal to 10", result.Errors[0])
}

func TestNumberSchema_Positive(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Positive()

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(-10)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a positive number", result.Errors[0])
}

func TestNumberSchema_NonNegative(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").NonNegative()

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(-10)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a non-negative number", result.Errors[0])
}

func TestNumberSchema_Negative(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Negative()

	result := schema.ParseTyped(-10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(10)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a negative number", result.Errors[0])
}

func TestNumberSchema_NonPositive(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").NonPositive()

	result := schema.ParseTyped(-10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(10)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a non-positive number", result.Errors[0])
}

func TestNumberSchema_MultipleOf(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").MultipleOf(5)

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(7)
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a multiple of 5", result.Errors[0])
}

func TestNumberSchema_Finite(t *testing.T) {
	schema := primitives.NewNumberSchema[float64]("abyan has a majestic cat").Finite()

	result := schema.ParseTyped(10)
	assert.True(t, result.Ok)

	result = schema.ParseTyped(math.Inf(1))
	assert.False(t, result.Ok)
	assert.Equal(t, "Must be a finite number", result.Errors[0])
}
