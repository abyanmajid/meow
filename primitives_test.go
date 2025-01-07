package meow

import (
	"testing"
	"time"
)

func TestStringValid(t *testing.T) {
	schema := String("schema")

	inputs := []any{"Mark", "123", "", "~!@#$%^&*()_+`,./;'"}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%s', expected no error, but got %v", in, err)
		}
	}
}

func TestStringInvalid(t *testing.T) {
	schema := String("schema")

	inputs := []any{nil, true, false, 123, time.Now()}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%s', expected an error, but got none instead.", in)
		}
	}
}
