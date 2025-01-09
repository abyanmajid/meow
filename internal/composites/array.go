package composites

import (
	"fmt"
	"reflect"

	core "github.com/abyanmajid/v/internal"
	"github.com/abyanmajid/v/internal/primitives"
)

type ArrayElementSchema[T primitives.Number] interface {
	primitives.StringSchema | primitives.NumberSchema[T] | primitives.BooleanSchema | primitives.DateSchema | primitives.NilSchema | primitives.AnySchema | primitives.NeverSchema
}

type ArraySchema[T any] struct {
	Schema *core.Schema[[]T]
	Inner  *core.Schema[T]
}

func NewArraySchema[T any](path string, inner *core.Schema[T]) *ArraySchema[T] {
	return &ArraySchema[T]{
		Schema: &core.Schema[[]T]{
			Path:  path,
			Rules: []core.Rule[[]T]{},
		},
		Inner: inner,
	}
}

func (s *ArraySchema[T]) Parse(value interface{}) *core.Result[[]T] {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		return s.Schema.NewErrorResult("Must be an array")
	}

	var parsedArray []T
	finalResult := s.Schema.NewSuccessResult()

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i).Interface()
		parsedValue, ok := element.(T)

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

	baseResult := s.Schema.ParseGeneric(parsedArray)
	if !baseResult.Success {
		finalResult.Success = false
		finalResult.Errors = append(finalResult.Errors, baseResult.Errors...)
	}

	finalResult.Value = parsedArray
	return finalResult
}

func (s *ArraySchema[T]) ParseTyped(value []T) *core.Result[[]T] {
	finalResult := s.Schema.NewSuccessResult()

	for i, v := range value {
		innerResult := s.Inner.ParseGeneric(v)
		if !innerResult.Success {
			errorMessage := fmt.Sprintf("Element at index %d: %s", i, innerResult.Errors)
			finalResult.Success = false
			finalResult.Errors = append(finalResult.Errors, errorMessage)
		}
	}

	baseResult := s.Schema.ParseGeneric(value)
	if !baseResult.Success {
		finalResult.Success = false
		finalResult.Errors = append(finalResult.Errors, baseResult.Errors...)
	}

	finalResult.Value = value
	return finalResult
}

func (s *ArraySchema[T]) Nonempty() *ArraySchema[T] {
	s.Schema.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) == 0 {
			return s.Schema.NewErrorResult("Array must not be empty")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Min(minLength int) *ArraySchema[T] {
	s.Schema.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) < minLength {
			errorMessage := fmt.Sprintf("Array must have at least %d elements", minLength)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Max(maxLength int) *ArraySchema[T] {
	s.Schema.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) > maxLength {
			errorMessage := fmt.Sprintf("Array must have at most %d elements", maxLength)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *ArraySchema[T]) Length(exactLength int) *ArraySchema[T] {
	s.Schema.AddRule(func(value []T) *core.Result[[]T] {
		if len(value) != exactLength {
			errorMessage := fmt.Sprintf("Array must have exactly %d elements", exactLength)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}
