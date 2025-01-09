package literals

import core "github.com/abyanmajid/v/internal"

type EnumSchema[T comparable] struct {
	Schema *core.Schema[T]
	Enums  map[T]struct{}
}

func NewEnumSchema[T comparable](path string, allowedValues []T) *EnumSchema[T] {
	enumMap := make(map[T]struct{}, len(allowedValues))
	for _, value := range allowedValues {
		enumMap[value] = struct{}{}
	}

	return &EnumSchema[T]{
		Schema: &core.Schema[T]{
			Path:  path,
			Rules: []core.Rule[T]{},
		},
		Enums: enumMap,
	}
}

func (s *EnumSchema[T]) Parse(value interface{}) *core.Result[T] {
	typedValue, ok := value.(T)
	if !ok {
		return s.Schema.NewErrorResult("Invalid type.")
	}

	if _, exists := s.Enums[typedValue]; !exists {
		return s.Schema.NewErrorResult("Value is not in the allowed enum set.")
	}

	return s.Schema.ParseGeneric(typedValue)
}
