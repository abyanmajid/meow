package core

type Result[T any] struct {
	Success bool
	Value   T
	Path    string
	Errors  []string
}

type Rule[T any] func(T) *Result[T]

type Coerce[T any] struct {
	Schema *Schema[T]
}

type Schema[T any] struct {
	Path  string
	Rules []Rule[T]
}

func (s *Schema[T]) AddRule(rule Rule[T]) {
	s.Rules = append(s.Rules, rule)
}

func (s *Schema[T]) NewSuccessResult() *Result[T] {
	return &Result[T]{
		Success: true,
		Path:    s.Path,
	}
}

func (s *Schema[T]) NewErrorResult(errorMessage string) *Result[T] {
	return &Result[T]{
		Success: false,
		Path:    s.Path,
		Errors:  []string{errorMessage},
	}
}
