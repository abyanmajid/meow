package meow

import (
	"fmt"
	"strconv"
	"strings"
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
	if intVal, err := strconv.Atoi(value); err == nil {
		return intVal, nil
	}

	if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
		return floatVal, nil
	}

	if strings.Contains(value, ".") && !strings.Contains(value, "e") {
		return value, nil
	}

	return nil, fmt.Errorf("unable to convert string '%s' to a known type", value)
}
