package v

import (
	core "github.com/abyanmajid/v/internal"
	"github.com/abyanmajid/v/internal/coercion"
	"github.com/abyanmajid/v/internal/composites"
	"github.com/abyanmajid/v/internal/literals"
	"github.com/abyanmajid/v/internal/primitives"
)

type Numeric primitives.Number

func String(path string) *primitives.StringSchema {
	return primitives.NewStringSchema(path)
}

func Float(path string) *primitives.NumberSchema[float64] {
	return primitives.NewNumberSchema[float64](path)
}

func Integer(path string) *primitives.NumberSchema[int] {
	return primitives.NewNumberSchema[int](path)
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

func Literal[T comparable](path string, literalValue T) *literals.LiteralSchema[T] {
	return literals.NewLiteralSchema(path, literalValue)
}

func Enum[T comparable](path string, allowedValues []T) *literals.EnumSchema[T] {
	return literals.NewEnumSchema(path, allowedValues)
}

func Array[T any](path string, innerSchema *core.Schema[T]) *composites.ArraySchema[T] {
	return composites.NewArraySchema[T](path, innerSchema)
}

type coercionExports struct {
	String  func(path string) *coercion.CoerceStringSchema
	Float   func(path string) *coercion.CoerceNumberSchema[float64]
	Integer func(path string) *coercion.CoerceNumberSchema[int]
	Boolean func(path string) *coercion.CoerceBooleanSchema
	Date    func(path string) *coercion.CoerceDateSchema
}

var Coerce = coercionExports{
	String: func(path string) *coercion.CoerceStringSchema {
		return coercion.NewCoerceStringSchema(path)
	},
	Float: func(path string) *coercion.CoerceNumberSchema[float64] {
		return coercion.NewCoerceNumberSchema[float64](path)
	},
	Integer: func(path string) *coercion.CoerceNumberSchema[int] {
		return coercion.NewCoerceNumberSchema[int](path)
	},
	Boolean: func(path string) *coercion.CoerceBooleanSchema {
		return coercion.NewCoerceBooleanSchema(path)
	},
	Date: func(path string) *coercion.CoerceDateSchema {
		return coercion.NewCoerceDateSchema(path)
	},
}
