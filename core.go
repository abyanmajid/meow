package meow

type MeowSchema[T any] interface {
	Validate(input T) error
}

type MeowBaseSchema[T any] struct {
	validateFunc func(input T) error
}

// Check the provided input against the validation function defined in the base schema (MeowBaseSchema).
// If a validation function is set, it will be used to validate the input and return an error if the validation fails.
// If no validation function is set, it will return nil.
//
// Parameters:
//
//	input T - The input value to be validated.
//
// Returns:
//
//	error - An error if the validation fails, otherwise nil.
func (s *MeowBaseSchema[T]) Validate(input T) error {
	if s.validateFunc != nil {
		return s.validateFunc(input)
	}

	return nil
}
