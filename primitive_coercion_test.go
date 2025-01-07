package meow

import "testing"

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
		nil:                "nil",
	}

	for in, _ := range testCases {
		err := schema.Parse(in)
		if err != nil {
			t.Errorf("For a valid input '%v', expected no error, but got %v", in, err)
		}
	}
}
