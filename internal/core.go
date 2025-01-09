package core

type Result[T any] struct {
	Ok     bool
	Value  T
	Path   string
	Errors []string
}

type Rule[T any] func(T) *Result[T]

type Schema[T any] struct {
	Path  string
	Rules []Rule[T]
}

type CoerceSchema[T any] struct {
	Inner Schema[T]
}

func (s *Schema[T]) AddRule(rule Rule[T]) {
	s.Rules = append(s.Rules, rule)
}

func (s *Schema[T]) NewSuccessResult() *Result[T] {
	return &Result[T]{
		Ok:   true,
		Path: s.Path,
	}
}

func (s *Schema[T]) NewErrorResult(errorMessage string) *Result[T] {
	return &Result[T]{
		Ok:     false,
		Path:   s.Path,
		Errors: []string{errorMessage},
	}
}

func (s *Schema[T]) ParseGeneric(value T) *Result[T] {
	finalResult := s.NewSuccessResult()
	for _, assertRule := range s.Rules {
		assertionResult := assertRule(value)
		if !assertionResult.Ok {
			finalResult.Ok = false
			finalResult.Errors = append(finalResult.Errors, assertionResult.Errors...)
		}
	}

	finalResult.Value = value

	return finalResult
}
