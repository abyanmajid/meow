package meow

import (
	"math"
	"testing"
	"time"
)

func TestCoerceStringValid(t *testing.T) {
	schema := Coerce.String("schema")

	testCases := map[any]string{
		"already a string": "already a string",
		"":                 "",
		1:                  "1",
		float64(0.5):       "0.5",
		float32(0.5):       "0.5",
		int16(1):           "1",
		int32(1):           "1",
		int64(1):           "1",
		true:               "true",
		false:              "false",
		nil:                "null",
	}

	for in, expected := range testCases {
		res, err := schema.Parse(in)

		if err != nil {
			t.Errorf("For input '%v', expected no error, but got %v", in, err)
			continue

		}

		if res != expected {
			t.Errorf("For input '%v', expected '%v', but got '%v'", in, expected, res)
		}
	}
}

func TestCoerceStringInvalid(t *testing.T) {
	schema := Coerce.String("schema")

	testCases := []any{
		struct{}{},
		[]int{1, 2, 3},
		map[string]int{},
	}

	for _, input := range testCases {
		_, err := schema.Parse(input)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none", input)
		}
	}
}
func TestCoerceIntegerValid(t *testing.T) {
	schema := Coerce.Integer("schema")

	testCases := map[any]int{
		1:            1,
		int8(1):      1,
		int16(1):     1,
		int32(1):     1,
		int64(1):     1,
		uint(1):      1,
		uint8(1):     1,
		uint16(1):    1,
		uint32(1):    1,
		uint64(1):    1,
		float32(1.5): 1,
		float64(1.5): 1,
		true:         1,
		false:        0,
		nil:          0,
	}

	for in, expected := range testCases {
		res, err := schema.Parse(in)

		if err != nil {
			t.Errorf("For input '%v', expected no error, but got %v", in, err)
			continue
		}

		if res != expected {
			t.Errorf("For input '%v', expected '%v', but got '%v'", in, expected, res)
		}
	}
}

func TestCoerceIntegerInvalid(t *testing.T) {
	schema := Coerce.Integer("schema")

	testCases := []any{
		"string",
		struct{}{},
		[]int{1, 2, 3},
		map[string]int{},
		math.NaN(),
		math.Inf(1),
		math.Inf(-1),
	}

	for _, input := range testCases {
		_, err := schema.Parse(input)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none", input)
		}
	}
}
func TestCoerceFloatValid(t *testing.T) {
	schema := Coerce.Float("schema")

	testCases := map[any]float64{
		1:            1.0,
		int8(1):      1.0,
		int16(1):     1.0,
		int32(1):     1.0,
		int64(1):     1.0,
		uint(1):      1.0,
		uint8(1):     1.0,
		uint16(1):    1.0,
		uint32(1):    1.0,
		uint64(1):    1.0,
		float32(1.5): 1.5,
		float64(1.5): 1.5,
		"1.5":        1.5,
		true:         1.0,
		false:        0.0,
		nil:          0.0,
	}

	for in, expected := range testCases {
		res, err := schema.Parse(in)

		if err != nil {
			t.Errorf("For input '%v', expected no error, but got %v", in, err)
			continue
		}

		if res != expected {
			t.Errorf("For input '%v', expected '%v', but got '%v'", in, expected, res)
		}
	}
}

func TestCoerceFloatInvalid(t *testing.T) {
	schema := Coerce.Float("schema")

	testCases := []any{
		"string",
		struct{}{},
		[]int{1, 2, 3},
		map[string]int{},
		math.NaN(),
		math.Inf(1),
		math.Inf(-1),
	}

	for _, input := range testCases {
		_, err := schema.Parse(input)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none", input)
		}
	}
}
func TestCoerceBooleanValid(t *testing.T) {
	schema := Coerce.Boolean("schema")

	testCases := map[any]bool{
		true:         true,
		false:        false,
		"true":       true,
		"false":      false,
		"1":          true,
		"0":          false,
		0:            false,
		1:            true,
		int8(1):      true,
		int8(0):      false,
		int16(1):     true,
		int16(0):     false,
		int32(1):     true,
		int32(0):     false,
		int64(1):     true,
		int64(0):     false,
		uint(1):      true,
		uint(0):      false,
		uint8(1):     true,
		uint8(0):     false,
		uint16(1):    true,
		uint16(0):    false,
		uint32(1):    true,
		uint32(0):    false,
		uint64(1):    true,
		uint64(0):    false,
		float32(1.0): true,
		float32(0.0): false,
		float64(1.0): true,
		float64(0.0): false,
		nil:          false,
	}

	for in, expected := range testCases {
		res, err := schema.Parse(in)

		if err != nil {
			t.Errorf("For input '%v', expected no error, but got %v", in, err)
			continue
		}

		if res != expected {
			t.Errorf("For input '%v', expected '%v', but got '%v'", in, expected, res)
		}
	}
}

func TestCoerceBooleanInvalid(t *testing.T) {
	schema := Coerce.Boolean("schema")

	testCases := []any{
		"not a bool",
		struct{}{},
		[]int{1, 2, 3},
		map[string]int{},
	}

	for _, input := range testCases {
		_, err := schema.Parse(input)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none", input)
		}
	}
}
func TestCoerceDateValid(t *testing.T) {
	schema := Coerce.Date("schema")

	testCases := map[any]time.Time{
		"2006-01-02":          time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		"01/02/2006":          time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		"2006-01-02 15:04:05": time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		"02/01/2006 15:04:05": time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC): time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		nil: {},
	}

	for in, expected := range testCases {
		res, err := schema.Parse(in)

		if err != nil {
			t.Errorf("For input '%v', expected no error, but got %v", in, err)
			continue
		}

		if !res.Equal(expected) {
			t.Errorf("For input '%v', expected '%v', but got '%v'", in, expected, res)
		}
	}
}

func TestCoerceDateInvalid(t *testing.T) {
	schema := Coerce.Date("schema")

	testCases := []any{
		"invalid date",
		struct{}{},
		[]int{1, 2, 3},
		map[string]int{},
	}

	for _, input := range testCases {
		_, err := schema.Parse(input)
		if err == nil {
			t.Errorf("For input '%v', expected an error, but got none", input)
		}
	}
}
