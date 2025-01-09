package primitives

import (
	"fmt"
	"time"

	core "github.com/abyanmajid/z/internal"
)

type DateSchema struct {
	Base *core.Schema[time.Time]
}

func NewDateSchema(path string) *DateSchema {
	return &DateSchema{
		Base: &core.Schema[time.Time]{
			Path:  path,
			Rules: []core.Rule[time.Time]{},
		},
	}
}

func (s *DateSchema) Parse(value interface{}) *core.Result[time.Time] {
	valueTime, isTime := value.(time.Time)
	if !isTime {
		return s.Base.NewErrorResult("Must be a string.")
	}

	return s.Base.ParseGeneric(valueTime)
}

func (s *DateSchema) ParseTyped(value time.Time) *core.Result[time.Time] {
	return s.Base.ParseGeneric(value)
}

func (s *DateSchema) Min(earliest time.Time) *DateSchema {
	s.Base.AddRule(func(value time.Time) *core.Result[time.Time] {
		if value.Before(earliest) {
			errorMessage := fmt.Sprintf("Must be later than or equal to %v", earliest)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *DateSchema) Max(latest time.Time) *DateSchema {
	s.Base.AddRule(func(value time.Time) *core.Result[time.Time] {
		if value.After(latest) {
			errorMessage := fmt.Sprintf("Must be earlier than or equal to %v", latest)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}
