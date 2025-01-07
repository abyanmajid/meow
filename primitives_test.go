package meow

import (
	"math"
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

func TestNumberValid(t *testing.T) {
	schema := Number("schema")

	inputs := []any{1, -1, math.Inf(1), math.Inf(-1), 0.5, float32(0.5), 0}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%s', expected no error, but got %v", in, err)
		}
	}
}

func TestNumberInvalid(t *testing.T) {
	schema := Number("schema")

	inputs := []any{"Hello", nil, true, false, math.NaN(), "123.45"}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%s', expected an error, but got none instead.", in)
		}
	}
}

func TestBooleanValid(t *testing.T) {
	schema := Boolean("schema")

	inputs := []any{true, false}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%s', expected no error, but got %v", in, err)
		}
	}
}

func TestBooleanInvalid(t *testing.T) {
	schema := Boolean("schema")
	inputs := []any{"Hello", 123, nil, math.NaN(), -0.5, float32(0.5)}
	for _, in := range inputs {
		err := schema.Parse(in)
		if err == in {
			t.Errorf("For an invalid input '%s', expected an error, but got none instead.", in)
		}
	}
}
func TestDateValid(t *testing.T) {
	schema := Date("schema")

	inputs := []any{time.Now(), time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestDateInvalid(t *testing.T) {
	schema := Date("schema")

	inputs := []any{"Hello", 123, nil, true, false, math.NaN(), time.Now().String()}

	for _, in := range inputs {
		err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
