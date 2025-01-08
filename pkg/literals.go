package meow

import (
	"errors"
	"fmt"
	"reflect"
)

func Literal[T any](path string, value T) *MeowSchema[T] {
	return &MeowSchema[T]{
		Parse: func(input any) *MeowResult[T] {
			if reflect.DeepEqual(input, value) {
				return &MeowResult[T]{
					Path:  path,
					Value: value,
				}
			}
			errMsg := fmt.Sprintf("input does not match the literal value '%v'", value)
			return &MeowResult[T]{
				Path:  path,
				Error: errors.New(errMsg),
			}
		},
	}
}
