package coercion

import (
	"fmt"
	"regexp"

	core "github.com/abyanmajid/z/internal"
	"github.com/abyanmajid/z/internal/primitives"
)

type CoerceStringSchema struct {
	Inner *primitives.StringSchema
}

func NewCoerceStringSchema(path string) *CoerceStringSchema {
	return &CoerceStringSchema{
		Inner: primitives.NewStringSchema(path),
	}
}

func (c *CoerceStringSchema) Parse(value interface{}) *core.Result[string] {
	coercedValue := fmt.Sprint(value)
	return c.ParseTyped(coercedValue)
}

func (c *CoerceStringSchema) ParseTyped(value string) *core.Result[string] {
	return c.Inner.ParseTyped(value)
}

func (c *CoerceStringSchema) Min(minLength int) *CoerceStringSchema {
	c.Inner.Min(minLength)
	return c
}

func (c *CoerceStringSchema) Max(maxLength int) *CoerceStringSchema {
	c.Inner.Max(maxLength)
	return c
}

func (c *CoerceStringSchema) Length(length int) *CoerceStringSchema {
	c.Inner.Length(length)
	return c
}

func (c *CoerceStringSchema) Email() *CoerceStringSchema {
	c.Inner.Email()
	return c
}

func (c *CoerceStringSchema) URL() *CoerceStringSchema {
	c.Inner.URL()
	return c
}

func (c *CoerceStringSchema) Regex(regex *regexp.Regexp) *CoerceStringSchema {
	c.Inner.Regex(regex)
	return c
}

func (c *CoerceStringSchema) Includes(substr string) *CoerceStringSchema {
	c.Inner.Includes(substr)
	return c
}

func (c *CoerceStringSchema) StartsWith(prefix string) *CoerceStringSchema {
	c.Inner.StartsWith(prefix)
	return c
}

func (c *CoerceStringSchema) EndsWith(suffix string) *CoerceStringSchema {
	c.Inner.EndsWith(suffix)
	return c
}

func (c *CoerceStringSchema) Date() *CoerceStringSchema {
	c.Inner.Date()
	return c
}

func (c *CoerceStringSchema) Time() *CoerceStringSchema {
	c.Inner.Time()
	return c
}

func (c *CoerceStringSchema) IP() *CoerceStringSchema {
	c.Inner.IP()
	return c
}

func (c *CoerceStringSchema) CIDR() *CoerceStringSchema {
	c.Inner.CIDR()
	return c
}
