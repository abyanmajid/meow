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

func (nc *NumberSchema) Parse(value interface{}) *core.Result[float64] {
	valueFloat64, isFloat64 := value.(float64)
	if !isFloat64 {
		return nc.Base.NewErrorResult("Must be a string.")
	}

	return nc.Base.ParseGeneric(valueFloat64)
}

func (nc *NumberSchema) ParseTyped(value float64) *core.Result[float64] {
	return nc.Base.ParseGeneric(value)
}

func (nc *NumberSchema) Gt(lowerBound float64) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than %v", lowerBound)
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Gte(lowerBound float64) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < lowerBound {
			errorMessage := fmt.Sprintf("Must be greater than or equal to %v", lowerBound)
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Lt(upperBound float64) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than %v", upperBound)
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Lte(upperBound float64) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > upperBound {
			errorMessage := fmt.Sprintf("Must be smaller than or equal to %v", upperBound)
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Int() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Floor(value) != value {
			return nc.Base.NewErrorResult("Must be an integer")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Positive() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= 0 {
			return nc.Base.NewErrorResult("Must be a positive number")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) NonNegative() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < 0 {
			return nc.Base.NewErrorResult("Must be a non-negative number")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Negative() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= 0 {
			return nc.Base.NewErrorResult("Must be a negative number")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) NonPositive() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > 0 {
			return nc.Base.NewErrorResult("Must be a non-positive number")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) MultipleOf(step float64) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Mod(value, step) != 0 {
			errorMessage := fmt.Sprintf("Must be a multiple of %v", step)
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Finite() *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.IsInf(value, 0) {
			return nc.Base.NewErrorResult("Must be a finite number")
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}
