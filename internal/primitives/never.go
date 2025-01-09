package primitives

import core "github.com/abyanmajid/z/internal"

type NeverSchema struct {
	Base *core.Schema[interface{}]
}

func NewNeverSchema(path string) *NeverSchema {
	return &NeverSchema{
		Base: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *NeverSchema) Parse(value interface{}) *core.Result[interface{}] {
	return s.Base.NewErrorResult("Value is not allowed.")
}
