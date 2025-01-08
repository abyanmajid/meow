package meow

import "fmt"

func Enum[T comparable](values ...T) *MeowSchema[T] {
	validValues := make(map[T]struct{})
	for _, v := range values {
		validValues[v] = struct{}{}
	}

	return &MeowSchema[T]{
		Parse: func(input any) *MeowResult[T] {
			value, ok := input.(T)
			if !ok {
				return &MeowResult[T]{
					Error: fmt.Errorf("invalid type; expected one of: %v", values),
				}
			}

			if _, exists := validValues[value]; !exists {
				return &MeowResult[T]{
					Error: fmt.Errorf("invalid value; expected one of: %v", values),
				}
			}

			return &MeowResult[T]{
				Error: nil,
				Value: value,
			}
		},
	}
}
