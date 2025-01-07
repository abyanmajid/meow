package meow

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

type MeowSchema struct {
	Parse func(input any) error
}

func String(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				errMsg := fmt.Sprintf("[%s]: '%s' is not a string.", varName, input)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Number(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			var floatInput float64
			switch v := input.(type) {
			case int:
				floatInput = float64(v)
			case float32:
				floatInput = float64(v)
			case float64:
				floatInput = v
			default:
				errMsg := fmt.Sprintf("[%s]: '%v' is not a number.", varName, input)
				return errors.New(errMsg)
			}

			if math.IsNaN(floatInput) {
				errMsg := fmt.Sprintf("[%s]: '%v' is NaN.", varName, input)
				return errors.New(errMsg)
			}

			return nil
		},
	}
}

func Boolean(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Bool {
				errMsg := fmt.Sprintf("[%s]: '%s' is not a string.", varName, input)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Date(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Struct || reflect.TypeOf(input) != reflect.TypeOf(time.Time{}) {
				errMsg := fmt.Sprintf("[%s]: '%v' is not a valid date.", varName, input)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Nil(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input != nil {
				errMsg := fmt.Sprintf("[%s]: '%v' is not nil.", varName, input)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Any(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil {
				errMsg := fmt.Sprintf("[%s]: input is nil.", varName)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Never(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			errMsg := fmt.Sprintf("[%s]: input is never allowed.", varName)
			return errors.New(errMsg)
		},
	}
}
