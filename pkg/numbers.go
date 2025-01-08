package meow

import (
	"fmt"
	"math"
)

func (s *MeowSchema[float64]) Int() *MeowSchema[int] {
	return &MeowSchema[int]{
		Parse: func(input any) *MeowResult[int] {
			numInput := convertToFloat(input)
			numFloored := math.Floor(numInput)
			if numInput != numFloored {
				return &MeowResult[int]{Error: fmt.Errorf("value must be an integer")}
			}
			numInt := int(numFloored)
			return &MeowResult[int]{Value: numInt}
		},
	}
}

func (s *MeowSchema[float64]) Gt(value float64) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput <= numValue {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be greater than %f", numValue)}
			}

			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}

			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Gte(value float64) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput < numValue {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be greater than or equal to %f", numValue)}
			}

			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}

			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Lt(value float64) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if numInput >= numValue {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be less than %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Lte(value float64) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(input)
			if numInput > numValue {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be less than or equal to %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Positive() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if numInput <= 0 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be positive")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) NonNegative() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if numInput < 0 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be non-negative")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Negative() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if numInput >= 0 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be negative")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) NonPositive() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if numInput > 0 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be non-positive")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) MultipleOf(value float64) *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			numValue := convertToFloat(value)
			if math.Mod(numInput, numValue) != 0 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be a multiple of %f", numValue)}
			}
			revertedVal, err := convertFromFloat[string](numValue)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Finite() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if math.IsInf(numInput, 0) {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be finite")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}

func (s *MeowSchema[float64]) Safe() *MeowSchema[float64] {
	return &MeowSchema[float64]{
		Parse: func(input any) *MeowResult[float64] {
			numInput := convertToFloat(input)
			if numInput < math.MinInt64 || numInput > math.MaxInt64 {
				return &MeowResult[float64]{Error: fmt.Errorf("value must be a safe integer")}
			}
			revertedVal, err := convertFromFloat[string](numInput)
			if err != nil {
				return &MeowResult[float64]{Error: fmt.Errorf("invalid input format: %d", err)}
			}
			return &MeowResult[float64]{Value: revertedVal.(float64)}
		},
	}
}
