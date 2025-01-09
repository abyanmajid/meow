package primitives

import core "github.com/abyanmajid/z/internal"

type AnySchema struct {
	Base *core.Schema[interface{}]
}

func NewAnySchema(path string) *AnySchema {
	return &AnySchema{
		Base: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *AnySchema) Parse(value interface{}) *core.Result[interface{}] {
	result := s.Base.NewSuccessResult()
	result.Value = value
	return result
}
