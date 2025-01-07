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

func String(varName string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) (string, error) {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				errMsg := fmt.Sprintf("[%s]: '%s' is not a string.", varName, input)
				return "", errors.New(errMsg)
			}
			return input.(string), nil
		},
	}
}

func Integer(varName string) *MeowSchema[int] {
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
				errMsg := fmt.Sprintf("[%s]: '%v' is not an integer.", varName, input)
				return res, errors.New(errMsg)
			}

			return res, nil
		},
	}
}

func Float(varName string) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) (float64, error) {
			var res float64

			switch v := input.(type) {
			case float32:
				res = float64(v)
			case float64:
				res = v
			default:
				errMsg := fmt.Sprintf("[%s]: '%v' is not a float.", varName, input)
				return res, errors.New(errMsg)
			}

			if math.IsNaN(res) {
				errMsg := fmt.Sprintf("[%s]: '%v' is NaN.", varName, input)
				return res, errors.New(errMsg)
			}

			return res, nil
		},
	}
}

func Boolean(varName string) *MeowSchema[bool] {
	return &MeowSchema[bool]{
		Parse: func(input any) (bool, error) {
			if input == nil {
				errMsg := fmt.Sprintf("[%s]: input is nil.", varName)
				return false, errors.New(errMsg)
			}

			switch v := input.(type) {
			case bool:
				return v, nil
			default:
				errMsg := fmt.Sprintf("[%s]: '%v' is not a boolean.", varName, input)
				return false, errors.New(errMsg)
			}
		},
	}
}

func Date(varName string) *MeowSchema[time.Time] {
	return &MeowSchema[time.Time]{
		Parse: func(input any) (time.Time, error) {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Struct || reflect.TypeOf(input) != reflect.TypeOf(time.Time{}) {
				errMsg := fmt.Sprintf("[%s]: '%v' is not a valid date.", varName, input)
				return time.Time{}, errors.New(errMsg)
			}
			return input.(time.Time), nil
		},
	}
}

func Nil(varName string) *MeowSchema[interface{}] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			if input != nil {
				errMsg := fmt.Sprintf("[%s]: '%v' is not nil.", varName, input)
				return nil, errors.New(errMsg)
			}
			return nil, nil
		},
	}
}

func Any(varName string) *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			if input == nil {
				errMsg := fmt.Sprintf("[%s]: input is nil.", varName)
				return nil, errors.New(errMsg)
			}
			return input, nil
		},
	}
}

func Never(varName string) *MeowSchema[any] {
	return &MeowSchema[any]{
		Parse: func(input any) (any, error) {
			errMsg := fmt.Sprintf("[%s]: input is never allowed.", varName)
			return nil, errors.New(errMsg)
		},
	}
}
