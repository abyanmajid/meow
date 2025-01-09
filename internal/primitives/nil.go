package primitives

import core "github.com/abyanmajid/z/internal"

type NilSchema struct {
	Base *core.Schema[interface{}]
}

func NewNilSchema(path string) *NilSchema {
	return &NilSchema{
		Base: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *NilSchema) Parse(value interface{}) *core.Result[interface{}] {
	if value != nil {
		return s.Base.NewErrorResult("Value must be nil.")
	}

	return s.Base.NewSuccessResult()
}
