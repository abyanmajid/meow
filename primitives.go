package meow

import (
	"errors"
	"reflect"
	"time"
)

type MeowStringSchema struct {
	MeowBaseSchema[string]
}

type MeowNumberSchema struct {
	MeowBaseSchema[float64]
}

type MeowBooleanSchema struct {
	MeowBaseSchema[bool]
}

type MeowDateSchema struct {
	MeowBaseSchema[time.Time]
}

type MeowNullSchema struct {
	MeowBaseSchema[interface{}]
}

type MeowAnySchema struct {
	MeowBaseSchema[interface{}]
}

type MeowNeverSchema struct {
	MeowBaseSchema[interface{}]
}

func String() *MeowStringSchema {
	return &MeowStringSchema{
		MeowBaseSchema[string]{
			validateFunc: func(input string) error {
				if reflect.TypeOf(input).Kind() != reflect.String {
					return errors.New("input is not a string")
				}
				return nil
			},
		},
	}
}

func Number() *MeowNumberSchema {
	return &MeowNumberSchema{
		MeowBaseSchema[float64]{
			validateFunc: func(input float64) error {
				if reflect.TypeOf(input).Kind() != reflect.Float64 {
					return errors.New("input is not a number")
				}
				return nil
			},
		},
	}
}

func Boolean() *MeowBooleanSchema {
	return &MeowBooleanSchema{
		MeowBaseSchema[bool]{
			validateFunc: func(input bool) error {
				if reflect.TypeOf(input).Kind() != reflect.Bool {
					return errors.New("input is not a boolean")
				}
				return nil
			},
		},
	}
}

func Date() *MeowDateSchema {
	return &MeowDateSchema{
		MeowBaseSchema[time.Time]{
			validateFunc: func(input time.Time) error {
				return nil
			},
		},
	}
}

func Nil() *MeowNullSchema {
	return &MeowNullSchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				if input != nil {
					return errors.New("input is neither nil")
				}
				return nil
			},
		},
	}
}

func Any() *MeowAnySchema {
	return &MeowAnySchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				return nil
			},
		},
	}
}

func Never() *MeowNeverSchema {
	return &MeowNeverSchema{
		MeowBaseSchema[interface{}]{
			validateFunc: func(input interface{}) error {
				return errors.New("input is never valid")
			},
		},
	}
}
