package coercion_test

import (
	"testing"
	"time"

	"github.com/abyanmajid/v/internal/coercion"
	"github.com/stretchr/testify/assert"
)

func TestNewCoerceDateSchema(t *testing.T) {
	path := "abyan has a majestic cat"
	schema := coercion.NewCoerceDateSchema(path)
	assert.NotNil(t, schema)
	assert.Equal(t, path, schema.Inner.Schema.Path)
}

func TestCoerceDateSchema_Parse(t *testing.T) {
	schema := coercion.NewCoerceDateSchema("abyan has a majestic cat")

	tests := []struct {
		input    interface{}
		expected time.Time
		isError  bool
	}{
		{time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), false},
		{"2021-01-01T00:00:00Z", time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), false},
		{int64(1609459200), time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), false},
		{1609459200, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), false},
		{1609459200.0, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), false},
		{"invalid date", time.Time{}, true},
		{true, time.Time{}, true},
	}

	for _, test := range tests {
		result := schema.Parse(test.input)

		if test.isError {
			assert.NotEmpty(t, result.Errors)
		} else {
			assert.Empty(t, result.Errors)

			expectedUTC := test.expected.UTC()
			actualUTC := result.Value.UTC()

			assert.Equal(t, expectedUTC, actualUTC)
		}
	}
}

func TestCoerceDateSchema_Min(t *testing.T) {
	schema := coercion.NewCoerceDateSchema("abyan has a majestic cat")
	minDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	schema.Min(minDate)
	resultValid := schema.Parse("2030-01-01T00:00:00Z")
	resultInvalid := schema.Parse("2010-01-01T00:00:00Z")
	assert.True(t, resultValid.Success)
	assert.False(t, resultInvalid.Success)
}

func TestCoerceDateSchema_Max(t *testing.T) {
	schema := coercion.NewCoerceDateSchema("abyan has a majestic cat")
	minDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	schema.Max(minDate)
	resultValid := schema.Parse("2010-01-01T00:00:00Z")
	resultInvalid := schema.Parse("2030-01-01T00:00:00Z")
	assert.True(t, resultValid.Success)
	assert.False(t, resultInvalid.Success)
}

func TestCoerceToInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
		ok       bool
	}{
		{int(123), 123, true},
		{int64(123), 123, true},
		{float64(123.0), 123, true},
		{"123", 0, false},
	}

	for _, test := range tests {
		result, ok := coercion.CoerceToInt64(test.input)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.ok, ok)
	}
}
