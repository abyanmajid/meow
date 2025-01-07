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

func (c *PrimitiveCoerce) String() *MeowSchema[string] {
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
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a string.", input, reflect.TypeOf(input))
				return "", errors.New(errMsg)
			}
			return str, nil
		},
	}
}

func (c *PrimitiveCoerce) Integer() *MeowSchema[int] {
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
					errMsg := fmt.Sprintf("'%v' is NaN or infinity, cannot convert to integer.", v)
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
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into an integer.", input, reflect.TypeOf(input))
				return 0, errors.New(errMsg)
			}
			return result, nil
		},
	}
}

func (c *PrimitiveCoerce) Float() *MeowSchema[float64] {
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
					errMsg := fmt.Sprintf("cannot parse '%v' into a float.", v)
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
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a float.", input, reflect.TypeOf(input))
				return 0, errors.New(errMsg)
			}

			if math.IsNaN(result) || math.IsInf(result, 0) {
				errMsg := fmt.Sprintf("'%v' is NaN or infinity, cannot convert to float.", result)
				return 0, errors.New(errMsg)
			}

			return result, nil
		},
	}
}

func (c *PrimitiveCoerce) Boolean() *MeowSchema[bool] {
	return &MeowSchema[bool]{
		Parse: func(input any) (bool, error) {
			var result bool
			switch v := input.(type) {
			case bool:
				result = v
			case string:
				switch v {
				case "true", "TRUE", "1":
					result = true
				case "false", "FALSE", "0":
					result = false
				default:
					parsed, err := strconv.ParseBool(v)
					if err != nil {
						errMsg := fmt.Sprintf("cannot parse '%v' into a bool.", v)
						return false, errors.New(errMsg)
					}
					result = parsed
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				if reflect.ValueOf(v).IsZero() {
					result = false
				} else {
					result = true
				}
			case float32, float64:
				if reflect.ValueOf(v).IsZero() {
					result = false
				} else {
					result = true
				}
			case nil:
				result = false
			default:
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a bool.", input, reflect.TypeOf(input))
				return false, errors.New(errMsg)
			}
			return result, nil
		},
	}
}

func (c *PrimitiveCoerce) Date() *MeowSchema[time.Time] {
	return &MeowSchema[time.Time]{
		Parse: func(input any) (time.Time, error) {
			var result time.Time
			switch v := input.(type) {
			case time.Time:
				result = v
			case string:
				layouts := []string{
					"2006-01-02",          // YYYY-MM-DD
					"01/02/2006",          // MM/DD/YYYY
					"2006-01-02 15:04:05", // YYYY-MM-DD HH:MM:SS
					"02/01/2006 15:04:05", // DD/MM/YYYY HH:MM:SS
				}
				var err error
				for _, layout := range layouts {
					result, err = time.Parse(layout, v)
					if err == nil {
						return result, nil
					}
				}
				errMsg := fmt.Sprintf("cannot parse '%v' into a date.", v)
				return result, errors.New(errMsg)
			case nil:
				result = time.Time{}
			default:
				errMsg := fmt.Sprintf("cannot coerce '%v' of type '%s' into a date.", input, reflect.TypeOf(input))
				return result, errors.New(errMsg)
			}
			return result, nil
		},
	}
}
