package coercion

import (
	"strings"

	core "github.com/abyanmajid/z/internal"
	"github.com/abyanmajid/z/internal/primitives"
)

type CoerceBooleanSchema struct {
	Inner *primitives.BooleanSchema
}

func NewCoerceBooleanSchema(path string) *CoerceBooleanSchema {
	return &CoerceBooleanSchema{
		Inner: primitives.NewBooleanSchema(path),
	}
}

func (c *CoerceBooleanSchema) Parse(value interface{}) *core.Result[bool] {
	var coercedValue bool
	switch v := value.(type) {
	case bool:
		coercedValue = v
	case string:
		v = strings.ToLower(v)
		if v == "true" {
			coercedValue = true
		} else if v == "false" {
			coercedValue = false
		} else {
			return c.Inner.Base.NewErrorResult("Must be a value that can be casted to a boolean")
		}
	case int:
		if v == 0 {
			coercedValue = false
		} else if v == 1 {
			coercedValue = true
		} else {
			return c.Inner.Base.NewErrorResult("Must be a value that can be casted to a boolean")
		}
	default:
		return c.Inner.Base.NewErrorResult("Must be a value that can be casted to a boolean")
	}

	return c.ParseTyped(coercedValue)
}

func (c *CoerceBooleanSchema) ParseTyped(value bool) *core.Result[bool] {
	return c.Inner.ParseTyped(value)
}
