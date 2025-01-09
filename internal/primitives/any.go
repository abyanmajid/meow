package primitives

import core "github.com/abyanmajid/v/internal"

type AnySchema struct {
	Schema *core.Schema[interface{}]
}

func NewAnySchema(path string) *AnySchema {
	return &AnySchema{
		Schema: &core.Schema[interface{}]{
			Path:  path,
			Rules: []core.Rule[interface{}]{},
		},
	}
}

func (s *AnySchema) Parse(value interface{}) *core.Result[interface{}] {
	result := s.Schema.NewSuccessResult()
	result.Value = value
	return result
}
