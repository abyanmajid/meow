package x

import (
	"github.com/abyanmajid/z/internal/coercion"
	"github.com/abyanmajid/z/internal/primitives"
)

func String(path string) *primitives.StringSchema {
	return primitives.NewStringSchema(path)
}

func Number(path string) *primitives.NumberSchema {
	return primitives.NewNumberSchema(path)
}

type CoerceSchema struct {
	String func(path string) *coercion.CoerceStringSchema
	Number func(path string) *coercion.CoerceNumberSchema
}

var Coerce = CoerceSchema{
	String: func(path string) *coercion.CoerceStringSchema {
		return coercion.NewCoerceStringSchema(path)
	},
	Number: func(path string) *coercion.CoerceNumberSchema {
		return coercion.NewCoerceNumberSchema(path)
	},
}
