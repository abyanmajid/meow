package meow

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

type PrimitiveCoerce struct{}

var Coerce = PrimitiveCoerce{}

func (c *PrimitiveCoerce) String(path string) *StringSchema[string] {
	return &StringSchema[string]{&Schema[string]{
		Parse: func(input any) *Result[string] {
			var str string
			switch input.(type) {
			case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
				str = input.(string)
			case nil:
				str = "null"
			default:
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a string.", input, reflect.TypeOf(input))
				return &Result[string]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			}
			return &Result[string]{
				Path:  path,
				Error: nil,
				Value: str,
			}
		},
	}}
}

func (c *PrimitiveCoerce) Number(path string) *NumberSchema[float64] {
	return &NumberSchema[float64]{&Schema[float64]{
		Parse: func(input any) *Result[float64] {
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
					errMsg := fmt.Sprintf("cannot parse '%v' into a float.", v)
					return &Result[float64]{
						Path:  path,
						Error: errors.New(errMsg),
					}
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
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a float.", input, reflect.TypeOf(input))
				return &Result[float64]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			}

			if math.IsNaN(result) || math.IsInf(result, 0) {
				errMsg := fmt.Sprintf("'%v' is NaN or infinity, cannot convert to float.", result)
				return &Result[float64]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			}

			return &Result[float64]{
				Path:  path,
				Value: result,
			}
		},
	}}
}

func (c *PrimitiveCoerce) Boolean(path string) *BooleanSchema[bool] {
	return &BooleanSchema[bool]{&Schema[bool]{
		Parse: func(input any) *Result[bool] {
			var res bool
			switch input := input.(type) {
			case bool:
				res = input
			case string:
				s := input
				switch s {
				case "true", "TRUE", "1":
					res = true
				case "false", "FALSE", "0":
					res = false
				default:
					parsed, err := strconv.ParseBool(s)
					if err != nil {
						errMsg := fmt.Sprintf("cannot parse '%v' into a bool.", s)
						return &Result[bool]{
							Path:  path,
							Error: errors.New(errMsg),
						}
					}
					res = parsed
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				n := input.(int)
				if reflect.ValueOf(n).IsZero() {
					res = false
				} else {
					res = true
				}
			case float32, float64:
				n := input.(float64)
				if reflect.ValueOf(n).IsZero() {
					res = false
				} else {
					res = true
				}
			case nil:
				res = false
			default:
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a bool.", input, reflect.TypeOf(input))
				return &Result[bool]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			}

			return &Result[bool]{
				Path:  path,
				Value: res,
			}
		},
	}}
}

func (c *PrimitiveCoerce) Date(path string) *DateSchema[time.Time] {
	return &DateSchema[time.Time]{&Schema[time.Time]{
		Parse: func(input any) *Result[time.Time] {
			var result time.Time
			switch input := input.(type) {
			case time.Time:
				result = input
			case string:
				s := input
				layouts := []string{
					"2006-01-02",          // YYYY-MM-DD
					"01/02/2006",          // MM/DD/YYYY
					"2006-01-02 15:04:05", // YYYY-MM-DD HH:MM:SS
					"02/01/2006 15:04:05", // DD/MM/YYYY HH:MM:SS
				}
				var err error
				for _, layout := range layouts {
					result, err = time.Parse(layout, s)
					if err == nil {
						return &Result[time.Time]{
							Path:  path,
							Value: result,
						}
					}
				}
				errMsg := fmt.Sprintf("cannot parse '%v' into a date.", s)
				return &Result[time.Time]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			case nil:
				result = time.Time{}
			default:
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a date.", input, reflect.TypeOf(input))
				return &Result[time.Time]{
					Path:  path,
					Error: errors.New(errMsg),
				}
			}
			return &Result[time.Time]{
				Path:  path,
				Value: result,
			}
		},
	}}
}
