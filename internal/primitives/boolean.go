package primitives

import core "github.com/abyanmajid/z/internal"

type BooleanSchema struct {
	Base *core.Schema[bool]
}

func NewBooleanSchema(path string) *BooleanSchema {
	return &BooleanSchema{
		Base: &core.Schema[bool]{
			Path:  path,
			Rules: []core.Rule[bool]{},
		},
	}
}

func (s *BooleanSchema) Parse(value interface{}) *core.Result[bool] {
	valueBool, isBool := value.(bool)
	if !isBool {
		return s.Base.NewErrorResult("Must be a boolean")
	}

	return s.Base.ParseGeneric(valueBool)
}

func (s *BooleanSchema) ParseTyped(value bool) *core.Result[bool] {
	return s.Base.ParseGeneric(value)
}
