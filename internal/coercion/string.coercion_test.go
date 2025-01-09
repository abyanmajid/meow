package coercion_test

import (
	"regexp"
	"testing"

	"github.com/abyanmajid/z/internal/coercion"
	"github.com/stretchr/testify/assert"
)

func TestNewCoerceStringSchema(t *testing.T) {
	path := "im sad"
	schema := coercion.NewCoerceStringSchema(path)
	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Inner.Base.Path)
}

func TestCoerceStringSchema_Parse(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	result := schema.Parse(123)
	assert.NotNil(t, result)
	assert.Equal(t, "123", result.Value)
}

func TestCoerceStringSchema_ParseTyped(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	result := schema.ParseTyped("test")
	assert.NotNil(t, result)
	assert.Equal(t, "test", result.Value)
}

func TestCoerceStringSchema_Min(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Min(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Max(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Max(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Length(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Length(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Email(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Email()
	result := schema.ParseTyped("test@example.com")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_URL(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.URL()
	result := schema.ParseTyped("https://example.com")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Regex(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	regex := regexp.MustCompile(`^[a-z]+$`)
	schema.Regex(regex)
	result := schema.ParseTyped("test")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Includes(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Includes("test")
	result := schema.ParseTyped("this is a test")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_StartsWith(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.StartsWith("test")
	result := schema.ParseTyped("test123")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_EndsWith(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.EndsWith("123")
	result := schema.ParseTyped("test123")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Date(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Date()
	result := schema.ParseTyped("2023-10-10")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Time(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.Time()
	result := schema.ParseTyped("10:10:10")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_IP(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.IP()
	result := schema.ParseTyped("192.168.0.1")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_CIDR(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("im sad")
	schema.CIDR()
	result := schema.ParseTyped("192.168.0.0/16")
	assert.True(t, result.Success)
}
