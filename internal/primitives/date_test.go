package primitives

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDateSchema(t *testing.T) {
	schema := NewDateSchema("abyan has a majestic cat")

	assert.NotNil(t, schema)
	assert.Equal(t, "abyan has a majestic cat", schema.Schema.Path)
	assert.Empty(t, schema.Schema.Rules)
}

func TestDateSchema_Parse(t *testing.T) {
	schema := NewDateSchema("abyan has a majestic cat")

	t.Run("valid time.Time", func(t *testing.T) {
		now := time.Now()
		result := schema.Parse(now)
		assert.True(t, result.Ok)
		assert.Equal(t, now, result.Value)
	})

	t.Run("invalid type", func(t *testing.T) {
		result := schema.Parse("invalid")
		assert.False(t, result.Ok)
		assert.Contains(t, result.Errors, "Must be a string.")
	})
}

func TestDateSchema_ParseTyped(t *testing.T) {
	schema := NewDateSchema("abyan has a majestic cat")
	now := time.Now()
	result := schema.ParseTyped(now)

	assert.True(t, result.Ok)
	assert.Equal(t, now, result.Value)
}

func TestDateSchema_Min(t *testing.T) {
	schema := NewDateSchema("abyan has a majestic cat")
	earliest := time.Now().Add(-time.Hour)
	schema.Min(earliest)

	t.Run("valid date", func(t *testing.T) {
		validDate := time.Now()
		result := schema.ParseTyped(validDate)
		assert.True(t, result.Ok)
	})

	t.Run("invalid date", func(t *testing.T) {
		invalidDate := time.Now().Add(-2 * time.Hour)
		result := schema.ParseTyped(invalidDate)
		assert.False(t, result.Ok)
		assert.Equal(t, fmt.Sprintf("Must be later than or equal to %v", earliest), result.Errors[0])
	})
}

func TestDateSchema_Max(t *testing.T) {
	schema := NewDateSchema("abyan has a majestic cat")
	latest := time.Now().Add(time.Hour)
	schema.Max(latest)

	t.Run("valid date", func(t *testing.T) {
		validDate := time.Now()
		result := schema.ParseTyped(validDate)
		assert.True(t, result.Ok)
	})

	t.Run("invalid date", func(t *testing.T) {
		invalidDate := time.Now().Add(2 * time.Hour)
		result := schema.ParseTyped(invalidDate)
		assert.False(t, result.Ok)
		assert.Equal(t, fmt.Sprintf("Must be earlier than or equal to %v", latest), result.Errors[0])
	})
}
