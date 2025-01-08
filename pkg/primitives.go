package meow

import (
	"math"
	"reflect"
	"time"
)

func String(path string) *StringSchema[string] {
	return &StringSchema[string]{&Schema[string]{
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
		Optional: false,
	}}
}

func Number(path string) *NumberSchema[float64] {
	return &NumberSchema[float64]{&Schema[float64]{
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
		Optional: false,
	}}
}

func Boolean(path string) *BooleanSchema[bool] {
	return &BooleanSchema[bool]{&Schema[bool]{
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
		Optional: false,
	}}
}

func Date(path string) *DateSchema[time.Time] {
	return &DateSchema[time.Time]{
		&Schema[time.Time]{
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
			Optional: false,
		}}
}

func Nil(path string) *NilSchema[any] {
	return &NilSchema[any]{
		&Schema[any]{
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
			Optional: false,
		},
	}
}

func Any(path string) *AnySchema[any] {
	return &AnySchema[any]{&Schema[any]{
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
		Optional: false,
	}}
}

func Never(path string) *NeverSchema[any] {
	return &NeverSchema[any]{
		&Schema[any]{
			Parse: func(input any) *Result[any] {
				return &Result[any]{
					Path:  path,
					Error: ErrInvalidNever,
				}
			},
			Optional: false,
		},
	}
}
