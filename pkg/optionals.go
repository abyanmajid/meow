package meow

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
				return &Result[*T]{
					Error: result.Error,
				}
			}

			return &Result[*T]{
				Value: &result.Value,
				Error: nil,
			}
		},
	}
}
