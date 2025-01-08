package meow

import (
	"errors"
	"fmt"
	"reflect"
)

func Literal[T any](path string, value T) *Schema[T] {
	return &Schema[T]{
		Parse: func(input any) *Result[T] {
			if reflect.DeepEqual(input, value) {
				return &Result[T]{
					Path:  path,
					Value: value,
				}
			}
			errMsg := fmt.Sprintf("input does not match the literal value '%v'", value)
			return &Result[T]{
				Path:  path,
				Error: errors.New(errMsg),
			}
		},
	}
}
