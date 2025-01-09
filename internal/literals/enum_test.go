package literals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnumSchema(t *testing.T) {
	allowedValues := []string{"value1", "value2", "value3"}
	enumSchema := NewEnumSchema("abyan has a majestic cat", allowedValues)

	assert.Equal(t, "abyan has a majestic cat", enumSchema.Schema.Path)
	assert.Len(t, enumSchema.Enums, len(allowedValues))
	for _, value := range allowedValues {
		_, exists := enumSchema.Enums[value]
		assert.True(t, exists)
	}
}

func TestEnumSchema_Parse(t *testing.T) {
	allowedValues := []string{"value1", "value2", "value3"}
	enumSchema := NewEnumSchema("abyan has a majestic cat", allowedValues)

	tests := []struct {
		input    interface{}
		expected string
	}{
		{"value1", ""},
		{"value4", "Value is not in the allowed enum set."},
		{123, "Invalid type."},
	}

	for _, test := range tests {
		result := enumSchema.Parse(test.input)
		if test.expected == "" {
			assert.True(t, result.Ok)
		} else {
			assert.False(t, result.Ok)
			assert.Contains(t, result.Errors, test.expected)
		}
	}
}
