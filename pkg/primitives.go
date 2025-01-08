package meow

import (
	"math"
	"reflect"
	"time"
)

type Result[T any] struct {
	Path  string
	Error error
	Value T
}

type Schema[T any] struct {
	Parse func(input any) *Result[T]
}

type MeowNonPrimitiveSchema[T any] struct {
	Parse func(input T) *Result[T]
}

func String(path string) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				return &Result[string]{
					Path:  path,
					Error: ErrInvalidString,
				}
			}
			return &Result[string]{
				Path:  path,
				Error: nil,
				Value: input.(string),
			}
		},
	}
}

func Number(path string) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			switch v := input.(type) {
			case float32, float64:
				if math.IsNaN(v.(float64)) {
					return &Result[float64]{
						Path:  path,
						Error: ErrInvalidFloat,
					}
				}
				return &Result[float64]{
					Path:  path,
					Error: nil,
					Value: v.(float64),
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				return &Result[float64]{
					Path:  path,
					Error: nil,
					Value: float64(reflect.ValueOf(v).Int()),
				}
			default:
				return &Result[float64]{
					Path:  path,
					Error: ErrInvalidFloat,
				}
			}
		},
	}
}

func Boolean(path string) *Schema[bool] {
	return &Schema[bool]{
		Parse: func(input any) *Result[bool] {
			switch input := input.(type) {
			case bool:
				return &Result[bool]{
					Path:  path,
					Error: nil,
					Value: input,
				}
			default:
				return &Result[bool]{
					Path:  path,
					Error: ErrInvalidBoolean,
				}
			}
		},
	}
}

func Date(path string) *Schema[time.Time] {
	return &Schema[time.Time]{
		Parse: func(input any) *Result[time.Time] {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Struct || reflect.TypeOf(input) != reflect.TypeOf(time.Time{}) {
				return &Result[time.Time]{
					Path:  path,
					Error: ErrInvalidBoolean,
				}
			}
			return &Result[time.Time]{
				Path:  path,
				Error: nil,
				Value: input.(time.Time),
			}
		},
	}
}

func Nil(path string) *Schema[any] {
	return &Schema[any]{
		Parse: func(input any) *Result[any] {
			if input != nil {
				return &Result[any]{
					Path:  path,
					Error: ErrInvalidNil,
				}
			}
			return &Result[any]{
				Path:  path,
				Error: nil,
				Value: nil,
			}
		},
	}
}

func Any(path string) *Schema[any] {
	return &Schema[any]{
		Parse: func(input any) *Result[any] {
			if input == nil {
				return &Result[any]{
					Path:  path,
					Error: ErrInvalidAny,
				}
			}
			return &Result[any]{
				Path:  path,
				Error: nil,
				Value: input,
			}
		},
	}
}

func Never(path string) *Schema[any] {
	return &Schema[any]{
		Parse: func(input any) *Result[any] {
			return &Result[any]{
				Path:  path,
				Error: ErrInvalidNever,
			}
		},
	}
}
