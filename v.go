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

func Boolean(path string) *primitives.BooleanSchema {
	return primitives.NewBooleanSchema(path)
}

func Date(path string) *primitives.DateSchema {
	return primitives.NewDateSchema(path)
}

func Nil(path string) *primitives.NilSchema {
	return primitives.NewNilSchema(path)
}

func Any(path string) *primitives.AnySchema {
	return primitives.NewAnySchema(path)
}

func Never(path string) *primitives.NeverSchema {
	return primitives.NewNeverSchema(path)
}

type coercionExports struct {
	String  func(path string) *coercion.CoerceStringSchema
	Number  func(path string) *coercion.CoerceNumberSchema
	Boolean func(path string) *coercion.CoerceBooleanSchema
	Date    func(path string) *coercion.CoerceDateSchema
}

var Coerce = coercionExports{
	String: func(path string) *coercion.CoerceStringSchema {
		return coercion.NewCoerceStringSchema(path)
	},
	Number: func(path string) *coercion.CoerceNumberSchema {
		return coercion.NewCoerceNumberSchema(path)
	},
	Boolean: func(path string) *coercion.CoerceBooleanSchema {
		return coercion.NewCoerceBooleanSchema(path)
	},
	Date: func(path string) *coercion.CoerceDateSchema {
		return coercion.NewCoerceDateSchema(path)
	},
}
