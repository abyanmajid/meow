package meow

import (
	"fmt"
)

func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	// Add more cases as needed
	default:
		return fmt.Sprintf("%v", v)
	}
}

func convertFromString[T any](value string) (any, error) {
	return value, nil
}
