package coercion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoerceBooleanSchema_Parse(t *testing.T) {
	schema := NewCoerceBooleanSchema("abyan has a majestic cat")

	tests := []struct {
		input    interface{}
		expected bool
		isError  bool
	}{
		{input: true, expected: true, isError: false},
		{input: false, expected: false, isError: false},
		{input: "true", expected: true, isError: false},
		{input: "false", expected: false, isError: false},
		{input: "TRUE", expected: true, isError: false},
		{input: "FALSE", expected: false, isError: false},
		{input: 1, expected: true, isError: false},
		{input: 0, expected: false, isError: false},
		{input: "invalid", expected: false, isError: true},
		{input: 2, expected: false, isError: true},
		{input: nil, expected: false, isError: true},
	}

	for _, test := range tests {
		result := schema.Parse(test.input)
		if test.isError {
			assert.False(t, result.Ok)
		} else {
			assert.True(t, result.Ok)
			assert.Equal(t, test.expected, result.Value)
		}
	}
}

func TestCoerceBooleanSchema_ParseTyped(t *testing.T) {
	schema := NewCoerceBooleanSchema("abyan has a majestic cat")

	tests := []struct {
		input    bool
		expected bool
	}{
		{input: true, expected: true},
		{input: false, expected: false},
	}

	for _, test := range tests {
		result := schema.ParseTyped(test.input)
		assert.True(t, result.Ok)
		assert.Equal(t, test.expected, result.Value)
	}
}
