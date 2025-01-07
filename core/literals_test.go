package meow

import (
	"testing"
)

func TestLiteral(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		input    any
		expected any
		wantErr  bool
	}{
		{
			name:     "int match",
			value:    42,
			input:    42,
			expected: 42,
			wantErr:  false,
		},
		{
			name:     "int mismatch",
			value:    42,
			input:    43,
			expected: 42,
			wantErr:  true,
		},
		{
			name:     "string match",
			value:    "hello",
			input:    "hello",
			expected: "hello",
			wantErr:  false,
		},
		{
			name:     "string mismatch",
			value:    "hello",
			input:    "world",
			expected: "hello",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema := Literal(tt.value)
			got, err := schema.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Literal().Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Literal().Parse() = %v, want %v", got, tt.expected)
			}
		})
	}
}
