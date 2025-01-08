package meow

import (
	"fmt"
	"reflect"
)

func Struct[T any](shape map[string]*Schema[any]) *StructSchema[T] {
	return &StructSchema[T]{
		&Schema[T]{
			Parse: func(input any) *Result[T] {
				typedInput, ok := input.(map[string]any)
				if !ok {
					return &Result[T]{Error: fmt.Errorf("input must be an object")}
				}

				result := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Elem()

				for key, schema := range shape {
					fieldValue, exists := typedInput[key]
					if !exists && !schema.Optional {
						if !exists {
							return &Result[T]{Error: fmt.Errorf("missing required key: %s", key)}
						}
					}
					if !exists && schema.Optional {
						continue
					}

					parseResult := schema.Parse(fieldValue)
					if parseResult.Error != nil {
						return &Result[T]{Error: fmt.Errorf("error in key %s: %v", key, parseResult.Error)}
					}

					result.FieldByName(key).Set(reflect.ValueOf(parseResult.Value))
				}

				return &Result[T]{Value: result.Interface().(T)}
			},
		},
	}
}

func Field[T any](schema *Schema[T]) *Schema[any] {
	return &Schema[any]{
		Parse: func(input any) *Result[any] {
			result := schema.Parse(input)
			if result.Error != nil {
				return &Result[any]{Error: result.Error}
			}
			return &Result[any]{Value: result.Value}
		},
		Optional: schema.Optional,
	}
}
