package literals

import core "github.com/abyanmajid/z/internal"

type EnumSchema[T comparable] struct {
	Base  *core.Schema[T]
	Enums map[T]struct{}
}

func NewEnumSchema[T comparable](path string, allowedValues []T) *EnumSchema[T] {
	enumMap := make(map[T]struct{}, len(allowedValues))
	for _, value := range allowedValues {
		enumMap[value] = struct{}{}
	}

	return &EnumSchema[T]{
		Base: &core.Schema[T]{
			Path:  path,
			Rules: []core.Rule[T]{},
		},
		Enums: enumMap,
	}
}

func (s *EnumSchema[T]) Parse(value interface{}) *core.Result[T] {
	typedValue, ok := value.(T)
	if !ok {
		return s.Base.NewErrorResult("Invalid type.")
	}

	if _, exists := s.Enums[typedValue]; !exists {
		return s.Base.NewErrorResult("Value is not in the allowed enum set.")
	}

	return s.Base.ParseGeneric(typedValue)
}
