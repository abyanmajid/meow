package coercion

import (
	"fmt"
	"time"

	core "github.com/abyanmajid/v/internal"
	"github.com/abyanmajid/v/internal/primitives"
)

type CoerceDateSchema struct {
	Inner *primitives.DateSchema
}

func NewCoerceDateSchema(path string) *CoerceDateSchema {
	return &CoerceDateSchema{
		Inner: primitives.NewDateSchema(path),
	}
}

func (c *CoerceDateSchema) Parse(value interface{}) *core.Result[time.Time] {
	var coercedValue time.Time
	switch v := value.(type) {
	case time.Time:
		coercedValue = v
	case string:
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return c.Inner.Schema.NewErrorResult(fmt.Sprintf("Must be a valid ISO 8601 date string, got: %v", v))
		}
		coercedValue = parsedTime
	case int, int64, float64:
		timestamp, ok := CoerceToInt64(v)
		if !ok {
			return c.Inner.Schema.NewErrorResult(fmt.Sprintf("Must be a valid Unix timestamp, got: %v", v))
		}
		coercedValue = time.Unix(timestamp, 0)
	default:
		return c.Inner.Schema.NewErrorResult("Must be a value that can be casted to a date")
	}

	return c.ParseTyped(coercedValue)
}

func (c *CoerceDateSchema) ParseTyped(value time.Time) *core.Result[time.Time] {
	return c.Inner.ParseTyped(value)
}

func (c *CoerceDateSchema) Min(earliest time.Time) *CoerceDateSchema {
	c.Inner.Min(earliest)
	return c
}

func (c *CoerceDateSchema) Max(latest time.Time) *CoerceDateSchema {
	c.Inner.Max(latest)
	return c
}

func CoerceToInt64(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int64:
		return v, true
	case float64:
		return int64(v), true
	default:
		return 0, false
	}
}
