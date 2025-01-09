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

func (csc *CoerceStringSchema) Parse(value interface{}) *core.Result[string] {
	coercedValue := fmt.Sprint(value)
	return csc.Inner.ParseTyped(coercedValue)
}

func (csc *CoerceStringSchema) ParseTyped(value string) *core.Result[string] {
	return csc.Inner.ParseTyped(value)
}

func (csc *CoerceStringSchema) Min(minLength int) *CoerceStringSchema {
	csc.Inner.Min(minLength)
	return csc
}

func (csc *CoerceStringSchema) Max(maxLength int) *CoerceStringSchema {
	csc.Inner.Max(maxLength)
	return csc
}

func (csc *CoerceStringSchema) Length(length int) *CoerceStringSchema {
	csc.Inner.Length(length)
	return csc
}

func (csc *CoerceStringSchema) Email() *CoerceStringSchema {
	csc.Inner.Email()
	return csc
}

func (csc *CoerceStringSchema) URL() *CoerceStringSchema {
	csc.Inner.URL()
	return csc
}

func (csc *CoerceStringSchema) Regex(regex *regexp.Regexp) *CoerceStringSchema {
	csc.Inner.Regex(regex)
	return csc
}

func (csc *CoerceStringSchema) Includes(substr string) *CoerceStringSchema {
	csc.Inner.Includes(substr)
	return csc
}

func (csc *CoerceStringSchema) StartsWith(prefix string) *CoerceStringSchema {
	csc.Inner.StartsWith(prefix)
	return csc
}

func (csc *CoerceStringSchema) EndsWith(suffix string) *CoerceStringSchema {
	csc.Inner.EndsWith(suffix)
	return csc
}

func (csc *CoerceStringSchema) Date() *CoerceStringSchema {
	csc.Inner.Date()
	return csc
}

func (csc *CoerceStringSchema) Time() *CoerceStringSchema {
	csc.Inner.Time()
	return csc
}

func (csc *CoerceStringSchema) IP() *CoerceStringSchema {
	csc.Inner.IP()
	return csc
}

func (csc *CoerceStringSchema) CIDR() *CoerceStringSchema {
	csc.Inner.CIDR()
	return csc
}
