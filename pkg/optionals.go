package meow

func Optional[T any](schema *MeowSchema[T]) *MeowSchema[*T] {
	return &MeowSchema[*T]{
		Parse: func(input any) *MeowResult[*T] {
			if input == nil {
				return &MeowResult[*T]{
					Value: nil,
					Error: nil,
				}
			}

			result := schema.Parse(input)
			if result.Error != nil {
				return &MeowResult[*T]{
					Error: result.Error,
				}
			}

			return &MeowResult[*T]{
				Value: &result.Value,
				Error: nil,
			}
		},
	}
}
