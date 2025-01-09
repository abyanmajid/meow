package coercion

import (
	"fmt"
	"strconv"

	core "github.com/abyanmajid/z/internal"
	"github.com/abyanmajid/z/internal/primitives"
)

type CoerceNumberSchema struct {
	Inner *primitives.NumberSchema
}

func NewCoerceNumberSchema(path string) *CoerceNumberSchema {
	return &CoerceNumberSchema{
		Inner: primitives.NewNumberSchema(path),
	}
}

func (cnc *CoerceNumberSchema) Parse(value interface{}) *core.Result[float64] {
	coercedValue := fmt.Sprint(value)

	parsedValue, err := strconv.ParseFloat(coercedValue, 64)
	if err != nil {
		return cnc.Inner.Base.NewErrorResult("Invalid number format")
	}

	return cnc.ParseTyped(parsedValue)
}

func (cnc *CoerceNumberSchema) ParseTyped(value float64) *core.Result[float64] {
	return cnc.Inner.ParseTyped(value)
}

func (cnc *CoerceNumberSchema) Gt(lowerBound float64) *CoerceNumberSchema {
	cnc.Inner.Gt(lowerBound)
	return cnc
}

func (cnc *CoerceNumberSchema) Gte(lowerBound float64) *CoerceNumberSchema {
	cnc.Inner.Gte(lowerBound)
	return cnc
}

func (cnc *CoerceNumberSchema) Lt(upperBound float64) *CoerceNumberSchema {
	cnc.Inner.Lt(upperBound)
	return cnc
}

func (cnc *CoerceNumberSchema) Lte(upperBound float64) *CoerceNumberSchema {
	cnc.Inner.Lte(upperBound)
	return cnc
}

func (cnc *CoerceNumberSchema) Int() *CoerceNumberSchema {
	cnc.Inner.Int()
	return cnc
}

func (cnc *CoerceNumberSchema) Positive() *CoerceNumberSchema {
	cnc.Inner.Positive()
	return cnc
}

func (cnc *CoerceNumberSchema) NonNegative() *CoerceNumberSchema {
	cnc.Inner.NonNegative()
	return cnc
}

func (cnc *CoerceNumberSchema) Negative() *CoerceNumberSchema {
	cnc.Inner.Negative()
	return cnc
}

func (cnc *CoerceNumberSchema) NonPositive() *CoerceNumberSchema {
	cnc.Inner.NonPositive()
	return cnc
}

func (cnc *CoerceNumberSchema) MultipleOf(step float64) *CoerceNumberSchema {
	cnc.Inner.MultipleOf(step)
	return cnc
}

func (cnc *CoerceNumberSchema) Finite() *CoerceNumberSchema {
	cnc.Inner.Finite()
	return cnc
}
