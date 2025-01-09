package coercion

import (
	"fmt"
	"strconv"

	core "github.com/abyanmajid/v/internal"
	"github.com/abyanmajid/v/internal/primitives"
)

type CoerceNumberSchema[T primitives.Number] struct {
	Inner *primitives.NumberSchema[T]
}

func NewCoerceNumberSchema[T primitives.Number](path string) *CoerceNumberSchema[T] {
	return &CoerceNumberSchema[T]{
		Inner: primitives.NewNumberSchema[T](path),
	}
}

func (c *CoerceNumberSchema[T]) Parse(value interface{}) *core.Result[T] {
	coercedValue := fmt.Sprint(value)

	parsedValue, err := strconv.ParseFloat(coercedValue, 64)
	if err != nil {
		return c.Inner.Schema.NewErrorResult("Must be a value that can be casted to a number")
	}

	return c.ParseTyped(T(parsedValue))
}

func (c *CoerceNumberSchema[T]) ParseTyped(value T) *core.Result[T] {
	return c.Inner.ParseTyped(value)
}

func (c *CoerceNumberSchema[T]) Gt(lowerBound T) *CoerceNumberSchema[T] {
	c.Inner.Gt(lowerBound)
	return c
}

func (c *CoerceNumberSchema[T]) Gte(lowerBound T) *CoerceNumberSchema[T] {
	c.Inner.Gte(lowerBound)
	return c
}

func (c *CoerceNumberSchema[T]) Lt(upperBound T) *CoerceNumberSchema[T] {
	c.Inner.Lt(upperBound)
	return c
}

func (c *CoerceNumberSchema[T]) Lte(upperBound T) *CoerceNumberSchema[T] {
	c.Inner.Lte(upperBound)
	return c
}

func (c *CoerceNumberSchema[T]) Positive() *CoerceNumberSchema[T] {
	c.Inner.Positive()
	return c
}

func (c *CoerceNumberSchema[T]) NonNegative() *CoerceNumberSchema[T] {
	c.Inner.NonNegative()
	return c
}

func (c *CoerceNumberSchema[T]) Negative() *CoerceNumberSchema[T] {
	c.Inner.Negative()
	return c
}

func (c *CoerceNumberSchema[T]) NonPositive() *CoerceNumberSchema[T] {
	c.Inner.NonPositive()
	return c
}

func (c *CoerceNumberSchema[T]) MultipleOf(step T) *CoerceNumberSchema[T] {
	c.Inner.MultipleOf(step)
	return c
}

func (c *CoerceNumberSchema[T]) Finite() *CoerceNumberSchema[T] {
	c.Inner.Finite()
	return c
}
