package primitives

import (
	"fmt"
	"time"

	core "github.com/abyanmajid/v/internal"
)

type DateSchema struct {
	Schema *core.Schema[time.Time]
}

func NewDateSchema(path string) *DateSchema {
	return &DateSchema{
		Schema: &core.Schema[time.Time]{
			Path:  path,
			Rules: []core.Rule[time.Time]{},
		},
	}
}

func (s *DateSchema) Parse(value interface{}) *core.Result[time.Time] {
	valueTime, isTime := value.(time.Time)
	if !isTime {
		return s.Schema.NewErrorResult("Must be a string.")
	}

	return s.Schema.ParseGeneric(valueTime)
}

func (s *DateSchema) ParseTyped(value time.Time) *core.Result[time.Time] {
	return s.Schema.ParseGeneric(value)
}

func (s *DateSchema) Min(earliest time.Time) *DateSchema {
	s.Schema.AddRule(func(value time.Time) *core.Result[time.Time] {
		if value.Before(earliest) {
			errorMessage := fmt.Sprintf("Must be later than or equal to %v", earliest)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *DateSchema) Max(latest time.Time) *DateSchema {
	s.Schema.AddRule(func(value time.Time) *core.Result[time.Time] {
		if value.After(latest) {
			errorMessage := fmt.Sprintf("Must be earlier than or equal to %v", latest)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}
