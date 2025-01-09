package literals

import (
	"fmt"

	core "github.com/abyanmajid/z/internal"
)

type LiteralSchema[T comparable] struct {
	Base  *core.Schema[T]
	Value T
}

func NewLiteralSchema[T comparable](path string, literalValue T) *LiteralSchema[T] {
	return &LiteralSchema[T]{
		Base: &core.Schema[T]{
			Path:  path,
			Rules: []core.Rule[T]{},
		},
		Value: literalValue,
	}
}

func (s *LiteralSchema[T]) Parse(value interface{}) *core.Result[T] {
	typedValue, ok := value.(T)
	if !ok {
		return s.Base.NewErrorResult("Invalid type.")
	}

	if typedValue != s.Value {
		return s.Base.NewErrorResult(fmt.Sprintf("Value must be %v.", s.Value))
	}

	return s.Base.ParseGeneric(typedValue)
}
