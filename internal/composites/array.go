package composites

import (
	"fmt"

	core "github.com/abyanmajid/z/internal"
)

type ArraySchema[T any] struct {
	Base  *core.Schema[[]T]
	Inner *core.Schema[T]
}

func NewArraySchema[T any](path string, inner *core.Schema[T]) *ArraySchema[T] {
	return &ArraySchema[T]{
		Base: &core.Schema[[]T]{
			Path:  path,
			Rules: []core.Rule[[]T]{},
		},
		Inner: inner,
	}
}

func (s *ArraySchema[T]) Parse(value interface{}) *core.Result[[]T] {
	values, isArray := value.([]interface{})
	if !isArray {
		return s.Base.NewErrorResult("Must be an array")
	}

	var parsedArray []T
	finalResult := s.Base.NewSuccessResult()

	for i, v := range values {
		parsedValue, ok := v.(T)
		if !ok {
			finalResult.Success = false
			finalResult.Errors = append(finalResult.Errors, fmt.Sprintf("Element at index %d must be of type %T", i, parsedValue))
			continue
		}

		innerResult := s.Inner.ParseGeneric(parsedValue)
		if !innerResult.Success {
			finalResult.Success = false
			finalResult.Errors = append(finalResult.Errors, innerResult.Errors...)
		} else {
			parsedArray = append(parsedArray, parsedValue)
		}
	}

	finalResult.Value = parsedArray
	return finalResult
}

func (s *ArraySchema[T]) ParseTyped(value []T) *core.Result[[]T] {
	finalResult := s.Base.NewSuccessResult()

	for i, v := range value {
		innerResult := s.Inner.ParseGeneric(v)
		if !innerResult.Success {
			errorMessage := fmt.Sprintf("Element at index %d: %s", i, innerResult.Errors)
			finalResult.Success = false
			finalResult.Errors = append(finalResult.Errors, errorMessage)
		}
	}

	finalResult.Value = value
	return finalResult
}

func (s *ArraySchema[T]) Nonempty() *ArraySchema[T] {
	s.Base.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) == 0 {
			return s.Base.NewErrorResult("Array must not be empty")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Min(minLength int) *ArraySchema[T] {
	s.Base.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) < minLength {
			errorMessage := fmt.Sprintf("Array must have at least %d elements", minLength)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Max(maxLength int) *ArraySchema[T] {
	s.Base.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) > maxLength {
			errorMessage := fmt.Sprintf("Array must have at most %d elements", maxLength)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Length(exactLength int) *ArraySchema[T] {
	s.Base.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) != exactLength {
			errorMessage := fmt.Sprintf("Array must have exactly %d elements", exactLength)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}
