package meow

import "fmt"

func Optional[T any](schema *Schema[T]) *Schema[*T] {
	return &Schema[*T]{
		Parse: func(input any) *Result[*T] {
			if input == nil {
				return &Result[*T]{
					Value: nil,
					Error: nil,
				}
			}

			result := schema.Parse(input)
			if result.Error != nil {
				fmt.Println("2")
				return &Result[*T]{
					Error: result.Error,
				}
			}

			return &Result[*T]{
				Value: &result.Value,
				Error: nil,
			}
		},
		Optional: true,
	}
}
