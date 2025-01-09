package primitives_test

import (
	"regexp"
	"testing"

	"github.com/abyanmajid/z/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewStringSchema(t *testing.T) {
	schema := primitives.NewStringSchema("im sad")
	assert.NotNil(t, schema)
	assert.Equal(t, "im sad", schema.Base.Path)
}

func TestStringSchema_Parse(t *testing.T) {
	schema := primitives.NewStringSchema("im sad")
	result := schema.Parse("test")
	assert.True(t, result.Success)

	result = schema.Parse(123)
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be a string.")
}

func TestStringSchema_Min(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Min(5, "Too short")
	result := schema.Parse("test")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Too short")

	result = schema.Parse("testing")
	assert.True(t, result.Success)
}

func TestStringSchema_Max(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Max(5, "Too long")
	result := schema.Parse("testing")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Too long")

	result = schema.Parse("test")
	assert.True(t, result.Success)
}

func TestStringSchema_Length(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Length(4, "Incorrect length")
	result := schema.Parse("test")
	assert.True(t, result.Success)

	result = schema.Parse("testing")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Incorrect length")
}

func TestStringSchema_Email(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Email("Invalid email")
	result := schema.Parse("test@example.com")
	assert.True(t, result.Success)

	result = schema.Parse("invalid-email")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid email")
}

func TestStringSchema_URL(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").URL("Invalid URL")
	result := schema.Parse("http://example.com")
	assert.True(t, result.Success)

	result = schema.Parse("invalid-url")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid URL")
}

func TestStringSchema_Regex(t *testing.T) {
	regex := regexp.MustCompile(`^[a-z]+$`)
	schema := primitives.NewStringSchema("im sad").Regex(regex, "Invalid format")
	result := schema.Parse("test")
	assert.True(t, result.Success)

	result = schema.Parse("Test123")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid format")
}

func TestStringSchema_Includes(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Includes("test", "Must include 'test'")
	result := schema.Parse("this is a test")
	assert.True(t, result.Success)

	result = schema.Parse("this is a sample")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must include 'test'")
}

func TestStringSchema_StartsWith(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").StartsWith("test", "Must start with 'test'")
	result := schema.Parse("test string")
	assert.True(t, result.Success)

	result = schema.Parse("string test")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must start with 'test'")
}

func TestStringSchema_EndsWith(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").EndsWith("test", "Must end with 'test'")
	result := schema.Parse("string test")
	assert.True(t, result.Success)

	result = schema.Parse("test string")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must end with 'test'")
}

func TestStringSchema_Date(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Date("Invalid date")
	result := schema.Parse("2023-10-10")
	assert.True(t, result.Success)

	result = schema.Parse("10-10-2023")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid date")
}

func TestStringSchema_Time(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Time("Invalid time")
	result := schema.Parse("15:04:05")
	assert.True(t, result.Success)

	result = schema.Parse("3:04 PM")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid time")
}

func TestStringSchema_IP(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").IP("Invalid IP")
	result := schema.Parse("192.168.0.1")
	assert.True(t, result.Success)

	result = schema.Parse("999.999.999.999")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid IP")
}

func TestStringSchema_CIDR(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").CIDR("Invalid CIDR")
	result := schema.Parse("192.168.0.0/24")
	assert.True(t, result.Success)

	result = schema.Parse("192.168.0.0/33")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Invalid CIDR")
}
