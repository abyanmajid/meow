package meow

import "fmt"

func Enum[T comparable](values ...T) *Schema[T] {
	validValues := make(map[T]struct{})
	for _, v := range values {
		validValues[v] = struct{}{}
	}

	return &Schema[T]{
		Parse: func(input any) *Result[T] {
			value, ok := input.(T)
			if !ok {
				return &Result[T]{
					Error: fmt.Errorf("invalid type; expected one of: %v", values),
				}
			}

			if _, exists := validValues[value]; !exists {
				return &Result[T]{
					Error: fmt.Errorf("invalid value; expected one of: %v", values),
				}
			}

			return &Result[T]{
				Error: nil,
				Value: value,
			}
		},
	}
}
