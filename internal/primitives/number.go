package primitives

import (
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

func (nc *NumberSchema) Gt(lowerBound float64, errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= lowerBound {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Gte(lowerBound float64, errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < lowerBound {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Lt(upperBound float64, errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= upperBound {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Lte(upperBound float64, errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > upperBound {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Int(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Floor(value) != value {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Positive(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value <= 0 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) NonNegative(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < 0 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Negative(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value >= 0 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) NonPositive(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value > 0 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) MultipleOf(step float64, errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.Mod(value, step) != 0 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Finite(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if math.IsInf(value, 0) {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}

func (nc *NumberSchema) Safe(errorMessage string) *NumberSchema {
	nc.Base.AddRule(func(value float64) *core.Result[float64] {
		if value < math.MinInt64 || value > math.MaxInt64 {
			return nc.Base.NewErrorResult(errorMessage)
		}
		return nc.Base.NewSuccessResult()
	})
	return nc
}
