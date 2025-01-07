package meow

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

type MeowSchema[T any] struct {
	Parse func(input any) (T, error)
}

func String() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) (string, error) {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				errMsg := fmt.Sprintf("'%s' is not a string", input)
				return "", errors.New(errMsg)
			}
			return input.(string), nil
		},
	}
}

func Integer() *MeowSchema[int] {
	return &MeowSchema[int]{
		Parse: func(input any) (int, error) {
			var res int

			switch v := input.(type) {
			case int:
				res = v
			case int8:
				res = int(v)
			case int16:
				res = int(v)
			case int32:
				res = int(v)
			case int64:
				res = int(v)
			case uint:
				res = int(v)
			case uint8:
				res = int(v)
			case uint16:
				res = int(v)
			case uint32:
				res = int(v)
			case uint64:
				res = int(v)
			default:
				errMsg := fmt.Sprintf("'%v' is not an integer", input)
				return res, errors.New(errMsg)
			}

			return res, nil
		},
	}
}

func Float() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) (float64, error) {
			var res float64

			switch v := input.(type) {
			case float32:
				res = float64(v)
			case float64:
				res = v
			default:
				errMsg := fmt.Sprintf("'%v' is not a float", input)
				return res, errors.New(errMsg)
			}

			if math.IsNaN(res) {
				errMsg := fmt.Sprintf("'%v' is NaN", input)
				return res, errors.New(errMsg)
			}

			return res, nil
		},
	}
}

func Boolean() *MeowSchema[bool] {
	return &MeowSchema[bool]{
		Parse: func(input any) (bool, error) {
			if input == nil {
				errMsg := "input is nil"
				return false, errors.New(errMsg)
			}

			switch v := input.(type) {
			case bool:
				return v, nil
			default:
				errMsg := fmt.Sprintf("'%v' is not a boolean", input)
				return false, errors.New(errMsg)
			}
		},
	}
}

func Date() *MeowSchema[time.Time] {
	return &MeowSchema[time.Time]{
		Parse: func(input any) (time.Time, error) {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Struct || reflect.TypeOf(input) != reflect.TypeOf(time.Time{}) {
				errMsg := fmt.Sprintf("'%v' is not a valid date", input)
				return time.Time{}, errors.New(errMsg)
			}
			return input.(time.Time), nil
		},
	}
}

func Nil() *MeowSchema[interface{}] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			if input != nil {
				errMsg := fmt.Sprintf("'%v' is not nil", input)
				return nil, errors.New(errMsg)
			}
			return nil, nil
		},
	}
}

func Any() *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			if input == nil {
				errMsg := "input is nil"
				return nil, errors.New(errMsg)
			}
			return input, nil
		},
	}
}

func Never() *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			errMsg := "input is never allowed"
			return nil, errors.New(errMsg)
		},
	}
}
