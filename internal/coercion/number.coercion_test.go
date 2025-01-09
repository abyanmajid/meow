package coercion_test

import (
	"testing"

	"github.com/abyanmajid/z/internal/coercion"
	"github.com/stretchr/testify/assert"
)

func TestNewCoerceNumberSchema(t *testing.T) {
	path := "im sad"
	schema := coercion.NewCoerceNumberSchema(path)
	assert.NotNil(t, schema)
	assert.NotNil(t, schema.Inner)
	assert.Equal(t, path, schema.Inner.Base.Path)
}

func TestCoerceNumberSchema_Parse(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	result := schema.Parse(123.45)
	assert.False(t, !result.Success)
	assert.Equal(t, 123.45, result.Value)
}

func TestCoerceNumberSchema_Gt(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad").Gt(10)
	result := schema.Parse(50)
	assert.True(t, result.Success)
	result = schema.Parse(-50)
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Gte(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Gte(10)
	result := schema.Parse("10")
	assert.True(t, result.Success)
	result = schema.Parse("5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Lt(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Lt(10)
	result := schema.Parse("5")
	assert.True(t, result.Success)
	result = schema.Parse("15")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Lte(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Lte(10)
	result := schema.Parse("10")
	assert.True(t, result.Success)
	result = schema.Parse("15")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Int(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Int()
	result := schema.Parse("10")
	assert.True(t, result.Success)
	result = schema.Parse("10.5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Positive(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Positive()
	result := schema.Parse("5")
	assert.True(t, result.Success)
	result = schema.Parse("-5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_NonNegative(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.NonNegative()
	result := schema.Parse("0")
	assert.True(t, result.Success)
	result = schema.Parse("-5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Negative(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Negative()
	result := schema.Parse("-5")
	assert.True(t, result.Success)
	result = schema.Parse("5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_NonPositive(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.NonPositive()
	result := schema.Parse("0")
	assert.True(t, result.Success)
	result = schema.Parse("5")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_MultipleOf(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.MultipleOf(5)
	result := schema.Parse("10")
	assert.True(t, result.Success)
	result = schema.Parse("7")
	assert.False(t, result.Success)
}

func TestCoerceNumberSchema_Finite(t *testing.T) {
	schema := coercion.NewCoerceNumberSchema("im sad")
	schema.Finite()
	result := schema.Parse("10")
	assert.True(t, result.Success)
	result = schema.Parse("Infinity")
	assert.False(t, result.Success)
}
