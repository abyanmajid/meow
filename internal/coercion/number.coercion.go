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

func (c *CoerceNumberSchema) Parse(value interface{}) *core.Result[float64] {
	coercedValue := fmt.Sprint(value)

	parsedValue, err := strconv.ParseFloat(coercedValue, 64)
	if err != nil {
		return c.Inner.Base.NewErrorResult("Must be a value that can be casted to a number")
	}

	return c.ParseTyped(parsedValue)
}

func (c *CoerceNumberSchema) ParseTyped(value float64) *core.Result[float64] {
	return c.Inner.ParseTyped(value)
}

func (c *CoerceNumberSchema) Gt(lowerBound float64) *CoerceNumberSchema {
	c.Inner.Gt(lowerBound)
	return c
}

func (c *CoerceNumberSchema) Gte(lowerBound float64) *CoerceNumberSchema {
	c.Inner.Gte(lowerBound)
	return c
}

func (c *CoerceNumberSchema) Lt(upperBound float64) *CoerceNumberSchema {
	c.Inner.Lt(upperBound)
	return c
}

func (c *CoerceNumberSchema) Lte(upperBound float64) *CoerceNumberSchema {
	c.Inner.Lte(upperBound)
	return c
}

func (c *CoerceNumberSchema) Int() *CoerceNumberSchema {
	c.Inner.Int()
	return c
}

func (c *CoerceNumberSchema) Positive() *CoerceNumberSchema {
	c.Inner.Positive()
	return c
}

func (c *CoerceNumberSchema) NonNegative() *CoerceNumberSchema {
	c.Inner.NonNegative()
	return c
}

func (c *CoerceNumberSchema) Negative() *CoerceNumberSchema {
	c.Inner.Negative()
	return c
}

func (c *CoerceNumberSchema) NonPositive() *CoerceNumberSchema {
	c.Inner.NonPositive()
	return c
}

func (c *CoerceNumberSchema) MultipleOf(step float64) *CoerceNumberSchema {
	c.Inner.MultipleOf(step)
	return c
}

func (c *CoerceNumberSchema) Finite() *CoerceNumberSchema {
	c.Inner.Finite()
	return c
}
