package meow

import (
	"errors"
	"fmt"
	"math"
	"reflect"
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

			if reflect.TypeOf(floatInput).Kind() != reflect.Float64 {
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
