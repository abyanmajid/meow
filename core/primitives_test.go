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
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%s', expected no error, but got %v", in, err)
		}
	}
}

func TestStringInvalid(t *testing.T) {
	schema := String("schema")

	inputs := []any{nil, true, false, 123, time.Now()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%s', expected an error, but got none instead.", in)
		}
	}
}

func TestFloatValid(t *testing.T) {
	schema := Float("schema")

	inputs := []any{float32(1.23), float64(4.56), float32(math.Pi), float64(math.E)}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestFloatInvalid(t *testing.T) {
	schema := Float("schema")

	inputs := []any{"Hello", 123, nil, true, false, struct{}{}, math.NaN()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}

func TestIntegerValid(t *testing.T) {
	schema := Integer("schema")

	inputs := []any{int(123), int8(123), int16(123), int32(123), int64(123), uint(123), uint8(123), uint16(123), uint32(123), uint64(123)}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestIntegerInvalid(t *testing.T) {
	schema := Integer("schema")

	inputs := []any{"Hello", 123.45, nil, true, false, struct{}{}, time.Now()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
func TestBooleanValid(t *testing.T) {
	schema := Boolean("schema")

	inputs := []any{true, false}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestBooleanInvalid(t *testing.T) {
	schema := Boolean("schema")

	inputs := []any{nil, "true", 1, 0, 123.45, struct{}{}, time.Now()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
func TestDateValid(t *testing.T) {
	schema := Date("schema")

	inputs := []any{time.Now(), time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestDateInvalid(t *testing.T) {
	schema := Date("schema")

	inputs := []any{nil, "2020-01-01", 123, true, false, struct{}{}}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
func TestNilValid(t *testing.T) {
	schema := Nil("schema")

	inputs := []any{nil}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestNilInvalid(t *testing.T) {
	schema := Nil("schema")

	inputs := []any{"Hello", 123, true, false, struct{}{}, time.Now()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
func TestAnyValid(t *testing.T) {
	schema := Any("schema")

	inputs := []any{"Hello", 123, true, false, 123.45, struct{}{}, time.Now()}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}

func TestAnyInvalid(t *testing.T) {
	schema := Any("schema")

	inputs := []any{nil}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For an invalid input '%v', expected an error, but got none instead.", in)
		}
	}
}
func TestNever(t *testing.T) {
	schema := Never("schema")

	inputs := []any{"Hello", 123, true, false, 123.45, struct{}{}, time.Now(), nil}

	for _, in := range inputs {
		_, err := schema.Parse(in)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none instead.", in)
		}
	}
}
