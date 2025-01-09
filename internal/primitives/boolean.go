package primitives

import core "github.com/abyanmajid/v/internal"

type BooleanSchema struct {
	Schema *core.Schema[bool]
}

func NewBooleanSchema(path string) *BooleanSchema {
	return &BooleanSchema{
		Schema: &core.Schema[bool]{
			Path:  path,
			Rules: []core.Rule[bool]{},
		},
	}
}

func (s *BooleanSchema) Parse(value interface{}) *core.Result[bool] {
	valueBool, isBool := value.(bool)
	if !isBool {
		return s.Schema.NewErrorResult("Must be a boolean")
	}

	return s.Schema.ParseGeneric(valueBool)
}

func (s *BooleanSchema) ParseTyped(value bool) *core.Result[bool] {
	return s.Schema.ParseGeneric(value)
}
