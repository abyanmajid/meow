package coercion_test

import (
	"testing"

	"github.com/abyanmajid/v/internal/coercion"
	"github.com/stretchr/testify/assert"
)

func TestNewCoerceNumberSchema(t *testing.T) {
	path := "abyan has a majestic cat"
	schema := coercion.NewCoerceNumberSchema[float64](path)
	assert.NotNil(t, schema)
	assert.NotNil(t, schema.Inner)
	assert.Equal(t, path, schema.Inner.Schema.Path)
}

func TestCoerceNumberSchema_Parse(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	result := schema.Parse(123.45)
	assert.False(t, !result.Ok)
	assert.Equal(t, 123.45, result.Value)
}

func TestCoerceNumberSchema_Gt(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat").Gt(10)
	result := schema.Parse(50)
	assert.True(t, result.Ok)
	result = schema.Parse(-50)
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Gte(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Gte(10)
	result := schema.Parse("10")
	assert.True(t, result.Ok)
	result = schema.Parse("5")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Lt(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Lt(10)
	result := schema.Parse("5")
	assert.True(t, result.Ok)
	result = schema.Parse("15")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Lte(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Lte(10)
	result := schema.Parse("10")
	assert.True(t, result.Ok)
	result = schema.Parse("15")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Positive(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Positive()
	result := schema.Parse("5")
	assert.True(t, result.Ok)
	result = schema.Parse("-5")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_NonNegative(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.NonNegative()
	result := schema.Parse("0")
	assert.True(t, result.Ok)
	result = schema.Parse("-5")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Negative(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Negative()
	result := schema.Parse("-5")
	assert.True(t, result.Ok)
	result = schema.Parse("5")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_NonPositive(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.NonPositive()
	result := schema.Parse("0")
	assert.True(t, result.Ok)
	result = schema.Parse("5")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_MultipleOf(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.MultipleOf(5)
	result := schema.Parse("10")
	assert.True(t, result.Ok)
	result = schema.Parse("7")
	assert.False(t, result.Ok)
}

func TestCoerceNumberSchema_Finite(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema[float64]("abyan has a majestic cat")
	schema.Finite()
	result := schema.Parse("10")
	assert.True(t, result.Ok)
	result = schema.Parse("Infinity")
	assert.False(t, result.Ok)
}
