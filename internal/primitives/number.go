package primitives

import (
	"fmt"
	"math"

	core "github.com/abyanmajid/z/internal"
)

type NumberSchema struct {
	Base *core.Schema[float64]
}

func NewNumberSchema(path string) *NumberSchema {
	return &NumberSchema{
		Base: &core.Schema[float64]{
			Path:  path,
			Rules: []core.Rule[float64]{},
		},
	}
}

func (s *NumberSchema) Parse(value interface{}) *core.Result[float64] {
	valueFloat64, isFloat64 := value.(float64)
	if !isFloat64 {
		return s.Base.NewErrorResult("Must be a string.")
	}

	return s.Base.ParseGeneric(valueFloat64)
}

func (s *NumberSchema) ParseTyped(value float64) *core.Result[float64] {
	return s.Base.ParseGeneric(value)
}

func (s *NumberSchema) Gt(lowerBound float64) *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than %v", lowerBound)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Gte(lowerBound float64) *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than or equal to %v", lowerBound)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Lt(upperBound float64) *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than %v", upperBound)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Lte(upperBound float64) *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than or equal to %v", upperBound)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Int() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Floor(value) != value {
			return s.Base.NewErrorResult("Must be an integer")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Positive() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= 0 {
			return s.Base.NewErrorResult("Must be a positive number")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) NonNegative() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < 0 {
			return s.Base.NewErrorResult("Must be a non-negative number")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Negative() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= 0 {
			return s.Base.NewErrorResult("Must be a negative number")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) NonPositive() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > 0 {
			return s.Base.NewErrorResult("Must be a non-positive number")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) MultipleOf(step float64) *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Mod(value, step) != 0 {
			errorMessage := fmt.Sprintf("Must be a multiple of %v", step)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *NumberSchema) Finite() *NumberSchema {
	s.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.IsInf(value, 0) {
			return s.Base.NewErrorResult("Must be a finite number")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}
