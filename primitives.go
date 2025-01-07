package meow

import (
	"errors"
	"fmt"
	"reflect"
)

type MeowSchema struct {
	Parse func(input any) error
}

func String(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.String {
				errMsg := fmt.Sprintf("'%s' is not a string.", varName)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}

func Number(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			if input == nil || reflect.TypeOf(input).Kind() != reflect.Float64 {
				errMsg := fmt.Sprintf("'%s' is not a number.", varName)
				return errors.New(errMsg)
			}
			return nil
		},
	}
}
