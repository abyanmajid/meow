package meow

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type PrimitiveCoerce struct{}

var Coerce = PrimitiveCoerce{}

func (c *PrimitiveCoerce) String(varName string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) (string, error) {
			var str string
			switch v := input.(type) {
			case string:
				str = v
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
				str = fmt.Sprintf("%v", v)
			case nil:
				str = "null"
			default:
				errMsg := fmt.Sprintf("[%s]: cannot coerce '%v' of type '%s' into a string.", varName, input, reflect.TypeOf(input))
				return "", errors.New(errMsg)
			}
			return str, nil
		},
	}
}

func (c *PrimitiveCoerce) Integer(varName string) *MeowSchema[int] {
	return &MeowSchema[int]{
		Parse: func(input any) (int, error) {
			var result int
			switch v := input.(type) {
			case int:
				result = v
			case int8:
				result = int(v)
			case int16:
				result = int(v)
			case int32:
				result = int(v)
			case int64:
				result = int(v)
			case uint:
				result = int(v)
			case uint8:
				result = int(v)
			case uint16:
				result = int(v)
			case uint32:
				result = int(v)
			case uint64:
				result = int(v)
			case float32:
				result = int(v)
			case float64:
				if math.IsNaN(v) || math.IsInf(v, 0) {
					errMsg := fmt.Sprintf("[%s]: '%v' is NaN or infinity, cannot convert to integer.", varName, v)
					return 0, errors.New(errMsg)
				}
				result = int(v)
			case bool:
				if v {
					result = 1
				} else {
					result = 0
				}
			case nil:
				result = 0
			default:
				errMsg := fmt.Sprintf("[%s]: cannot coerce '%v' of type '%s' into an integer.", varName, input, reflect.TypeOf(input))
				return 0, errors.New(errMsg)
			}
			return result, nil
		},
	}
}

func (c *PrimitiveCoerce) Float(varName string) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) (float64, error) {
			var result float64
			switch v := input.(type) {
			case float32:
				result = float64(v)
			case float64:
				result = v
			case int:
				result = float64(v)
			case int8:
				result = float64(v)
			case int16:
				result = float64(v)
			case int32:
				result = float64(v)
			case int64:
				result = float64(v)
			case uint:
				result = float64(v)
			case uint8:
				result = float64(v)
			case uint16:
				result = float64(v)
			case uint32:
				result = float64(v)
			case uint64:
				result = float64(v)
			case string:
				parsed, err := strconv.ParseFloat(v, 64)
				if err != nil {
					errMsg := fmt.Sprintf("[%s]: cannot parse '%v' into a float.", varName, v)
					return 0, errors.New(errMsg)
				}
				result = parsed
			case bool:
				if v {
					result = 1.0
				} else {
					result = 0.0
				}
			case nil:
				result = 0.0
			default:
				errMsg := fmt.Sprintf("[%s]: cannot coerce '%v' of type '%s' into a float.", varName, input, reflect.TypeOf(input))
				return 0, errors.New(errMsg)
			}

			// Handle NaN and Infinity
			if math.IsNaN(result) || math.IsInf(result, 0) {
				errMsg := fmt.Sprintf("[%s]: '%v' is NaN or infinity, cannot convert to float.", varName, result)
				return 0, errors.New(errMsg)
			}

			return result, nil
		},
	}
}
