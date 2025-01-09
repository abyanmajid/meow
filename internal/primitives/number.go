package primitives

import (
	"fmt"
	"math"

	core "github.com/abyanmajid/v/internal"
)

type Number interface {
	~float64 | ~float32 | ~int | ~int32 | ~int64
}

type NumberSchema[T Number] struct {
	Schema *core.Schema[T]
}

func NewNumberSchema[T Number](path string) *NumberSchema[T] {
	return &NumberSchema[T]{
		Schema: &core.Schema[T]{
			Path:  path,
			Rules: []core.Rule[T]{},
		},
	}
}

func (s *NumberSchema[T]) Parse(value interface{}) *core.Result[T] {
	valueT, isT := value.(T)
	if !isT {
		return s.Schema.NewErrorResult("Must be a number.")
	}

	return s.Schema.ParseGeneric(valueT)
}

func (s *NumberSchema[T]) ParseTyped(value T) *core.Result[T] {
	return s.Schema.ParseGeneric(value)
}

func (s *NumberSchema[T]) Gt(lowerBound T) *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value <= lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than %v", lowerBound)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Gte(lowerBound T) *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value < lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than or equal to %v", lowerBound)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Lt(upperBound T) *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value >= upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than %v", upperBound)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Lte(upperBound T) *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value > upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than or equal to %v", upperBound)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Positive() *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value <= 0 {
			return s.Schema.NewErrorResult("Must be a positive number")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) NonNegative() *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value < 0 {
			return s.Schema.NewErrorResult("Must be a non-negative number")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Negative() *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value >= 0 {
			return s.Schema.NewErrorResult("Must be a negative number")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) NonPositive() *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if value > 0 {
			return s.Schema.NewErrorResult("Must be a non-positive number")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) MultipleOf(step T) *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if math.Mod(float64(value), float64(step)) != 0 {
			errorMessage := fmt.Sprintf("Must be a multiple of %v", step)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema[T]) Finite() *NumberSchema[T] {
	s.Schema.AddRule(func(value T) *core.Result[T] {
		if math.IsInf(float64(value), 0) {
			return s.Schema.NewErrorResult("Must be a finite number")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}
