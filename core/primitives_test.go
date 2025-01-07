package meow

import (
	"errors"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected *MeowResult[string]
	}{
		{
			name:  "valid string",
			input: "hello",
			expected: &MeowResult[string]{
				Path:  "testPath",
				Error: nil,
				Value: "hello",
			},
		},
		{
			name:  "nil input",
			input: nil,
			expected: &MeowResult[string]{
				Path:  "testPath",
				Error: errors.New("not a valid string"),
			},
		},
		{
			name:  "non-string input",
			input: 123,
			expected: &MeowResult[string]{
				Path:  "testPath",
				Error: errors.New("not a valid string"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema := String("testPath")
			result := schema.Parse(tt.input)
			if result.Path != tt.expected.Path || result.Error != nil && result.Error.Error() != tt.expected.Error.Error() || result.Value != tt.expected.Value {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
