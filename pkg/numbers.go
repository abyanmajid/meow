package meow

import (
	"fmt"
	"math"
)

func (s *Schema[float64]) Int() *Schema[int] {
	return &Schema[int]{
		Parse: func(input any) *Result[int] {
			numInput := convertToFloat(input)
			numFloored := math.Floor(numInput)
			if numInput != numFloored {
				return &Result[int]{Error: fmt.Errorf("value must be an integer")}
			}
			numInt := int(numFloored)
			return &Result[int]{Value: numInt}
		},
	}
}

func (s *Schema[float64]) Gt(value float64) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput <= numValue {
				return &Result[float64]{Error: fmt.Errorf("value must be greater than %f", numValue)}
			}

			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}

			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Gte(value float64) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput < numValue {
				return &Result[float64]{Error: fmt.Errorf("value must be greater than or equal to %f", numValue)}
			}

			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}

			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Lt(value float64) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput >= numValue {
				return &Result[float64]{Error: fmt.Errorf("value must be less than %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Lte(value float64) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(input)
			if numInput > numValue {
				return &Result[float64]{Error: fmt.Errorf("value must be less than or equal to %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Positive() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if numInput <= 0 {
				return &Result[float64]{Error: fmt.Errorf("value must be positive")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) NonNegative() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if numInput < 0 {
				return &Result[float64]{Error: fmt.Errorf("value must be non-negative")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Negative() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if numInput >= 0 {
				return &Result[float64]{Error: fmt.Errorf("value must be negative")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) NonPositive() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if numInput > 0 {
				return &Result[float64]{Error: fmt.Errorf("value must be non-positive")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) MultipleOf(value float64) *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if math.Mod(numInput, numValue) != 0 {
				return &Result[float64]{Error: fmt.Errorf("value must be a multiple of %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Finite() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if math.IsInf(numInput, 0) {
				return &Result[float64]{Error: fmt.Errorf("value must be finite")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *Schema[float64]) Safe() *Schema[float64] {
	return &Schema[float64]{
		Parse: func(input any) *Result[float64] {
			numInput := convertToFloat(input)
			if numInput < math.MinInt64 || numInput > math.MaxInt64 {
				return &Result[float64]{Error: fmt.Errorf("value must be a safe integer")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &Result[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &Result[float64]{Value: revertedVal.(float64)}
		},
	}
}
