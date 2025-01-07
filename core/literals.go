package meow

import (
	"fmt"
	"reflect"
)

func Literal[T any](value T) *MeowSchema[T] {
	return &MeowSchema[T]{
		Parse: func(input any) (T, error) {
			if reflect.DeepEqual(input, value) {
				return value, nil
			}
			return value, fmt.Errorf("input '%v' does not match the literal value '%v'", input, value)
		},
	}
}
