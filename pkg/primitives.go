package meow

import (
	"math"
	"reflect"
	"time"
)

type MeowResult[T any] struct {
	Path  string
	Error error
	Value T
}

type MeowSchema[T any] struct {
	Parse func(input any) *MeowResult[T]
}

type MeowNonPrimitiveSchema[T any] struct {
	Parse func(input T) *MeowResult[T]
}

func String(path string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				return &MeowResult[string]{
					Path:  path,
					Error: ErrInvalidString,
				}
			}
			return &MeowResult[string]{
				Path:  path,
				Error: nil,
				Value: input.(string),
			}
		},
	}
}

func Number(path string) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			switch input.(type) {
			case float32, float64:
				if math.IsNaN(input.(float64)) {
					return &MeowResult[float64]{
						Path:  path,
						Error: ErrInvalidFloat,
					}
				}

				return &MeowResult[float64]{
					Path:  path,
					Error: nil,
					Value: input.(float64),
				}
			default:
				return &MeowResult[float64]{
					Path:  path,
					Error: ErrInvalidFloat,
				}
			}
		},
	}
}

func Boolean(path string) *MeowSchema[bool] {
	return &MeowSchema[bool]{
		Parse: func(input any) *MeowResult[bool] {
			switch input := input.(type) {
			case bool:
				return &MeowResult[bool]{
					Path:  path,
					Error: nil,
					Value: input,
				}
			default:
				return &MeowResult[bool]{
					Path:  path,
					Error: ErrInvalidBoolean,
				}
			}
		},
	}
}

func Date(path string) *MeowSchema[time.Time] {
	return &MeowSchema[time.Time]{
		Parse: func(input any) *MeowResult[time.Time] {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Struct || reflect.TypeOf(input) != reflect.TypeOf(time.Time{}) {
				return &MeowResult[time.Time]{
					Path:  path,
					Error: ErrInvalidBoolean,
				}
			}
			return &MeowResult[time.Time]{
				Path:  path,
				Error: nil,
				Value: input.(time.Time),
			}
		},
	}
}

func Nil(path string) *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) *MeowResult[any] {
			if input != nil {
				return &MeowResult[any]{
					Path:  path,
					Error: ErrInvalidNil,
				}
			}
			return &MeowResult[any]{
				Path:  path,
				Error: nil,
				Value: nil,
			}
		},
	}
}

func Any(path string) *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) *MeowResult[any] {
			if input == nil {
				return &MeowResult[any]{
					Path:  path,
					Error: ErrInvalidAny,
				}
			}
			return &MeowResult[any]{
				Path:  path,
				Error: nil,
				Value: input,
			}
		},
	}
}

func Never(path string) *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) *MeowResult[any] {
			return &MeowResult[any]{
				Path:  path,
				Error: ErrInvalidNever,
			}
		},
	}
}
