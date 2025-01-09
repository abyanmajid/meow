package coercion_test

import (
	"regexp"
	"testing"

	"github.com/abyanmajid/v/internal/coercion"
	"github.com/stretchr/testify/assert"
)

func TestNewCoerceStringSchema(t *testing.T) {
	path := "abyan has a majestic cat"
	schema := coercion.NewCoerceStringSchema(path)
	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Inner.Schema.Path)
}

func TestCoerceStringSchema_Parse(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	result := schema.Parse(123)
	assert.NotNil(t, result)
	assert.Equal(t, "123", result.Value)
}

func TestCoerceStringSchema_ParseTyped(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	result := schema.ParseTyped("test")
	assert.NotNil(t, result)
	assert.Equal(t, "test", result.Value)
}

func TestCoerceStringSchema_Min(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Min(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Max(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Max(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Length(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Length(5)
	result := schema.ParseTyped("12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Email(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Email()
	result := schema.ParseTyped("test@example.com")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_URL(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.URL()
	result := schema.ParseTyped("https://example.com")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Regex(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	regex := regexp.MustCompile(`^[a-z]+$`)
	schema.Regex(regex)
	result := schema.ParseTyped("test")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Includes(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Includes("test")
	result := schema.ParseTyped("this is a test")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_StartsWith(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.StartsWith("test")
	result := schema.ParseTyped("test123")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_EndsWith(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.EndsWith("123")
	result := schema.ParseTyped("test123")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Date(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Date()
	result := schema.ParseTyped("2023-10-10")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_Time(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.Time()
	result := schema.ParseTyped("10:10:10")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_IP(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.IP()
	result := schema.ParseTyped("192.168.0.1")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_CIDR(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.CIDR()
	result := schema.ParseTyped("192.168.0.0/16")
	assert.True(t, result.Success)
}
func TestCoerceStringSchema_UUID(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.UUID()
	result := schema.ParseTyped("123e4567-e89b-12d3-a456-426614174000")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_NanoID(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.NanoID()
	result := schema.ParseTyped("1234567890abcdef12345")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_CUID(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.CUID()
	result := schema.ParseTyped("c123456789012345678901234")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_CUID2(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.CUID2()
	result := schema.ParseTyped("c1234567890abcdef")
	assert.True(t, result.Success)
}

func TestCoerceStringSchema_ULID(t *testing.T) {
	schema := coercion.NewCoerceStringSchema("abyan has a majestic cat")
	schema.ULID()
	result := schema.ParseTyped("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	assert.True(t, result.Success)
}
