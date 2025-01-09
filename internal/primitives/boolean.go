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

func (nc *BooleanSchema) Parse(value interface{}) *core.Result[bool] {
	valueBool, isBool := value.(bool)
	if !isBool {
		return nc.Base.NewErrorResult("Must be a boolean")
	}

	return nc.Base.ParseGeneric(valueBool)
}

func (nc *BooleanSchema) ParseTyped(value bool) *core.Result[bool] {
	return nc.Base.ParseGeneric(value)
}
