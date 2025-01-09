package primitives

import core "github.com/abyanmajid/v/internal"

type NilSchema struct {
	Schema *core.Schema[interface{}]
}

func NewNilSchema(path string) *NilSchema {
	return &NilSchema{
		Schema: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *NilSchema) Parse(value interface{}) *core.Result[interface{}] {
	if value != nil {
		return s.Schema.NewErrorResult("Value must be nil.")
	}

	return s.Schema.NewSuccessResult()
}
