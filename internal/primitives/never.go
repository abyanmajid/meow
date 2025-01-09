package primitives

import core "github.com/abyanmajid/v/internal"

type NeverSchema struct {
	Schema *core.Schema[interface{}]
}

func NewNeverSchema(path string) *NeverSchema {
	return &NeverSchema{
		Schema: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *NeverSchema) Parse(value interface{}) *core.Result[interface{}] {
	return s.Schema.NewErrorResult("Value is not allowed.")
}
