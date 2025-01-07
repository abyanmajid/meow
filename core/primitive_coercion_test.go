package meow

import (
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
