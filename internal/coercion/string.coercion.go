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

func (csc *CoerceStringSchema) Min(minLength int, errorMessage string) *CoerceStringSchema {
	csc.Inner.Min(minLength, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Max(maxLength int, errorMessage string) *CoerceStringSchema {
	csc.Inner.Max(maxLength, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Length(length int, errorMessage string) *CoerceStringSchema {
	csc.Inner.Length(length, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Email(errorMessage string) *CoerceStringSchema {
	csc.Inner.Email(errorMessage)
	return csc
}

func (csc *CoerceStringSchema) URL(errorMessage string) *CoerceStringSchema {
	csc.Inner.URL(errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Regex(regex *regexp.Regexp, errorMessage string) *CoerceStringSchema {
	csc.Inner.Regex(regex, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Includes(substr, errorMessage string) *CoerceStringSchema {
	csc.Inner.Includes(substr, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) StartsWith(prefix, errorMessage string) *CoerceStringSchema {
	csc.Inner.StartsWith(prefix, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) EndsWith(suffix, errorMessage string) *CoerceStringSchema {
	csc.Inner.EndsWith(suffix, errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Date(errorMessage string) *CoerceStringSchema {
	csc.Inner.Date(errorMessage)
	return csc
}

func (csc *CoerceStringSchema) Time(errorMessage string) *CoerceStringSchema {
	csc.Inner.Time(errorMessage)
	return csc
}

func (csc *CoerceStringSchema) IP(errorMessage string) *CoerceStringSchema {
	csc.Inner.IP(errorMessage)
	return csc
}

func (csc *CoerceStringSchema) CIDR(errorMessage string) *CoerceStringSchema {
	csc.Inner.CIDR(errorMessage)
	return csc
}
