package meow

import (
	"errors"
	"reflect"
	"time"
)

// MeowStringSchema represents a schema for validating string inputs.
type MeowStringSchema struct {
	MeowBaseSchema[string]
}

// MeowNumberSchema represents a schema for validating number inputs (float64).
type MeowNumberSchema struct {
	MeowBaseSchema[float64]
}

// MeowBooleanSchema represents a schema for validating boolean inputs.
type MeowBooleanSchema struct {
	MeowBaseSchema[bool]
}

// MeowDateSchema represents a schema for validating date inputs (time.Time).
type MeowDateSchema struct {
	MeowBaseSchema[time.Time]
}

// MeowNullSchema represents a schema for validating nil inputs.
type MeowNullSchema struct {
	MeowBaseSchema[interface{}]
}

// MeowAnySchema represents a schema that validates any input without restrictions.
type MeowAnySchema struct {
	MeowBaseSchema[interface{}]
}

// MeowNeverSchema represents a schema that always fails validation.
type MeowNeverSchema struct {
	MeowBaseSchema[interface{}]
}

// String creates a schema for validating string inputs.
func String() *MeowStringSchema {
	return &MeowStringSchema{
		MeowBaseSchema[string]{
			validateFunc: func(input string) error {
				// Check if the input is a string
				if reflect.TypeOf(input).Kind() != reflect.String {
					return errors.New("input is not a string")
				}
				return nil
			},
		},
	}
}

// Number creates a schema for validating number inputs (float64).
func Number() *MeowNumberSchema {
	return &MeowNumberSchema{
		MeowBaseSchema[float64]{
			validateFunc: func(input float64) error {
				// Check if the input is a float64
				if reflect.TypeOf(input).Kind() != reflect.Float64 {
					return errors.New("input is not a number")
				}
				return nil
			},
		},
	}
}

// Boolean creates a schema for validating boolean inputs.
func Boolean() *MeowBooleanSchema {
	return &MeowBooleanSchema{
		MeowBaseSchema[bool]{
			validateFunc: func(input bool) error {
				// Check if the input is a boolean
				if reflect.TypeOf(input).Kind() != reflect.Bool {
					return errors.New("input is not a boolean")
				}
				return nil
			},
		},
	}
}

// Date creates a schema for validating date inputs (time.Time).
func Date() *MeowDateSchema {
	return &MeowDateSchema{
		MeowBaseSchema[time.Time]{
			validateFunc: func(input time.Time) error {
				// Since the input is already a time.Time, this is always valid
				return nil
			},
		},
	}
}

// Nil creates a schema for validating nil inputs.
func Nil() *MeowNullSchema {
	return &MeowNullSchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				// Check if the input is nil
				if input != nil {
					return errors.New("input is neither nil")
				}
				return nil
			},
		},
	}
}

// Any creates a schema that validates any input without restrictions.
func Any() *MeowAnySchema {
	return &MeowAnySchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				// Always valid
				return nil
			},
		},
	}
}

// Never creates a schema that always fails validation.
func Never() *MeowNeverSchema {
	return &MeowNeverSchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				// Always fails validation
				return errors.New("input is never valid")
			},
		},
	}
}
