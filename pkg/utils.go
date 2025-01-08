package meow

import (
	"fmt"
	"strconv"
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

func convertToFloat(input any) float64 {
	switch v := input.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return f
	default:
		return 0
	}
}

func convertFromFloat[T any](value float64) (any, error) {
	return value, nil
}
