package meow

import (
	"errors"
	"fmt"
	"reflect"
)

type PrimitiveCoerce struct{}

var Coerce = PrimitiveCoerce{}

func (c *PrimitiveCoerce) String(varName string) *MeowSchema {
	return &MeowSchema{
		Parse: func(input any) error {
			var str string
			switch v := input.(type) {
			case string:
				str = v
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
				str = fmt.Sprintf("%v", v)
			case nil:
				str = "null"
			default:
				errMsg := fmt.Sprintf("[%s]: cannot coerce '%s' of type '%s' into a string.", varName, input, reflect.TypeOf(input))
				return errors.New(errMsg)
			}
			if reflect.TypeOf(str).Kind() != reflect.String {
				return fmt.Errorf("input is not a string: %v", input)
			}
			return nil
		},
	}
}
