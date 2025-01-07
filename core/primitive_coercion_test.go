package meow

import (
	"math"
	"testing"
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
