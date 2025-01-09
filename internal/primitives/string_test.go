package primitives_test

import (
	"regexp"
	"testing"

	"github.com/abyanmajid/v/internal/primitives"
	"github.com/stretchr/testify/assert"
)

func TestNewStringSchema(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat")
	assert.NotNil(t, schema)
	assert.Equal(t, "abyan has a majestic cat", schema.Schema.Path)
}

func TestStringSchema_Parse(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat")
	result := schema.Parse("test")
	assert.True(t, result.Ok)

	result = schema.Parse(123)
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a string.")
}

func TestStringSchema_Min(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Min(5)
	result := schema.Parse("test")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be longer than 5 characters in length")

	result = schema.Parse("testing")
	assert.True(t, result.Ok)
}

func TestStringSchema_Max(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Max(5)
	result := schema.Parse("testing")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be shorter than 5 characters in length")

	result = schema.Parse("test")
	assert.True(t, result.Ok)
}

func TestStringSchema_Length(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Length(4)
	result := schema.Parse("test")
	assert.True(t, result.Ok)

	result = schema.Parse("testing")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be exactly 4 characters long")
}

func TestStringSchema_Email(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Email()
	result := schema.Parse("test@example.com")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-email")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid email address")
}

func TestStringSchema_URL(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").URL()
	result := schema.Parse("http://example.com")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-url")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid URL")
}

func TestStringSchema_Regex(t *testing.T) {
	regex := regexp.MustCompile(`^[a-z]+$`)
	schema := primitives.NewStringSchema("abyan has a majestic cat").Regex(regex)
	result := schema.Parse("test")
	assert.True(t, result.Ok)

	result = schema.Parse("Test123")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must match the required pattern")
}

func TestStringSchema_Includes(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Includes("test")
	result := schema.Parse("this is a test")
	assert.True(t, result.Ok)

	result = schema.Parse("this is a sample")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must include 'test'")
}

func TestStringSchema_StartsWith(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").StartsWith("test")
	result := schema.Parse("test string")
	assert.True(t, result.Ok)

	result = schema.Parse("string test")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must start with 'test'")
}

func TestStringSchema_EndsWith(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").EndsWith("test")
	result := schema.Parse("string test")
	assert.True(t, result.Ok)

	result = schema.Parse("test string")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must end with 'test'")
}

func TestStringSchema_Date(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Date()
	result := schema.Parse("2023-10-10")
	assert.True(t, result.Ok)

	result = schema.Parse("10-10-2023")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must follow a valid date format")
}

func TestStringSchema_Time(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").Time()
	result := schema.Parse("15:04:05")
	assert.True(t, result.Ok)

	result = schema.Parse("3:04 PM")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must follow a valid time format")
}

func TestStringSchema_IP(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").IP()
	result := schema.Parse("192.168.0.1")
	assert.True(t, result.Ok)

	result = schema.Parse("999.999.999.999")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid IP address")
}

func TestStringSchema_CIDR(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").CIDR()
	result := schema.Parse("192.168.0.0/24")
	assert.True(t, result.Ok)

	result = schema.Parse("192.168.0.0/33")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be of valid CIDR notation")
}
func TestStringSchema_UUID(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").UUID()
	result := schema.Parse("123e4567-e89b-12d3-a456-426614174000")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-uuid")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid UUID")
}

func TestStringSchema_NanoID(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").NanoID()
	result := schema.Parse("1234567890abcdef12345")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-nanoid")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid NanoID")
}

func TestStringSchema_CUID(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").CUID()
	result := schema.Parse("c123456789012345678901234")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-cuid")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid CUID")
}

func TestStringSchema_CUID2(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").CUID2()
	result := schema.Parse("cuid2example")
	assert.True(t, result.Ok)

	result = schema.Parse("InvalidCUID2")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid CUID2")
}

func TestStringSchema_ULID(t *testing.T) {
	schema := primitives.NewStringSchema("abyan has a majestic cat").ULID()
	result := schema.Parse("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	assert.True(t, result.Ok)

	result = schema.Parse("invalid-ulid")
	assert.False(t, result.Ok)
	assert.Contains(t, result.Errors, "Must be a valid ULID")
}
