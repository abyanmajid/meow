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
	schema := primitives.NewStringSchema("im sad").Min(5)
	result := schema.Parse("test")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be longer than 5 characters in length")

	result = schema.Parse("testing")
	assert.True(t, result.Success)
}

func TestStringSchema_Max(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Max(5)
	result := schema.Parse("testing")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be shorter than 5 characters in length")

	result = schema.Parse("test")
	assert.True(t, result.Success)
}

func TestStringSchema_Length(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Length(4)
	result := schema.Parse("test")
	assert.True(t, result.Success)

	result = schema.Parse("testing")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be exactly 4 characters long")
}

func TestStringSchema_Email(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Email()
	result := schema.Parse("test@example.com")
	assert.True(t, result.Success)

	result = schema.Parse("invalid-email")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be a valid email address")
}

func TestStringSchema_URL(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").URL()
	result := schema.Parse("http://example.com")
	assert.True(t, result.Success)

	result = schema.Parse("invalid-url")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be a valid URL")
}

func TestStringSchema_Regex(t *testing.T) {
	regex := regexp.MustCompile(`^[a-z]+$`)
	schema := primitives.NewStringSchema("im sad").Regex(regex)
	result := schema.Parse("test")
	assert.True(t, result.Success)

	result = schema.Parse("Test123")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must match the required pattern")
}

func TestStringSchema_Includes(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Includes("test")
	result := schema.Parse("this is a test")
	assert.True(t, result.Success)

	result = schema.Parse("this is a sample")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must include 'test'")
}

func TestStringSchema_StartsWith(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").StartsWith("test")
	result := schema.Parse("test string")
	assert.True(t, result.Success)

	result = schema.Parse("string test")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must start with 'test'")
}

func TestStringSchema_EndsWith(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").EndsWith("test")
	result := schema.Parse("string test")
	assert.True(t, result.Success)

	result = schema.Parse("test string")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must end with 'test'")
}

func TestStringSchema_Date(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Date()
	result := schema.Parse("2023-10-10")
	assert.True(t, result.Success)

	result = schema.Parse("10-10-2023")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must follow a valid date format")
}

func TestStringSchema_Time(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").Time()
	result := schema.Parse("15:04:05")
	assert.True(t, result.Success)

	result = schema.Parse("3:04 PM")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must follow a valid time format")
}

func TestStringSchema_IP(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").IP()
	result := schema.Parse("192.168.0.1")
	assert.True(t, result.Success)

	result = schema.Parse("999.999.999.999")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be a valid IP address")
}

func TestStringSchema_CIDR(t *testing.T) {
	schema := primitives.NewStringSchema("im sad").CIDR()
	result := schema.Parse("192.168.0.0/24")
	assert.True(t, result.Success)

	result = schema.Parse("192.168.0.0/33")
	assert.False(t, result.Success)
	assert.Contains(t, result.Errors, "Must be of valid CIDR notation")
}
