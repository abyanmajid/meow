package meow

import (
	"fmt"
	"reflect"
)

func Struct[T any](shape map[string]any) *Schema[T] {
	return &Schema[T]{
		Parse: func(input any) *Result[T] {
			typedInput, ok := input.(map[string]any)
			if !ok {
				return &Result[T]{Error: fmt.Errorf("input must be an object")}
			}

			result := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Elem()

			for key, schema := range shape {
				fieldValue, exists := typedInput[key]
				if !exists {
					return &Result[T]{Error: fmt.Errorf("missing required key: %s", key)}
				}

				schemaInstance, ok := schema.(*Schema[any])
				if !ok {
					return &Result[T]{Error: fmt.Errorf("invalid schema type for key: %s", key)}
				}

				fieldResult := schemaInstance.Parse(fieldValue)
				if fieldResult.Error != nil {
					return &Result[T]{Error: fmt.Errorf("error in key %s: %v", key, fieldResult.Error)}
				}

				result.FieldByName(key).Set(reflect.ValueOf(fieldResult.Value))
			}

			return &Result[T]{Value: result.Interface().(T)}
		},
	}
}
