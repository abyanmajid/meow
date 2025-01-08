package meow

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

func (s *StringSchema[string]) Datetime(offset bool, precision int) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			layout := "2006-01-02T15:04:05Z07:00" // ISO 8601 format with timezone
			if precision > 0 {
				layout = "2006-01-02T15:04:05." + strings.Repeat("0", precision) + "Z07:00"
			}
			_, err := time.Parse(layout, strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid datetime format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid datetime format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Date() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			_, err := time.Parse("2006-01-02", strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid date format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid date format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Time() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			layout := "15:04:05"
			_, err := time.Parse(layout, strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid time format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid time format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Ip() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			if net.ParseIP(strInput) == nil {
				return &Result[string]{Error: fmt.Errorf("invalid IP format: %s", strInput)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid IP format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Cidr() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			_, _, err := net.ParseCIDR(strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid CIDR format: %s", strInput)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid CIDR format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Min(min int) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			if len(strInput) < min {
				return &Result[string]{Error: fmt.Errorf("must be %d or more characters long", min)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Max(max int) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			if len(strInput) > max {
				return &Result[string]{Error: fmt.Errorf("must be %d or fewer characters long", max)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Length(exact int) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			if len(strInput) != exact {
				return &Result[string]{Error: fmt.Errorf("must be exactly %d characters long", exact)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Email() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)

			re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
			if !re.MatchString(strInput) {
				return &Result[string]{Error: fmt.Errorf("invalid email address")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid email format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Url() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)

			re := `^(http|https):\/\/[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
			match, err := regexp.MatchString(re, strInput)
			if err != nil || !match {
				return &Result[string]{Error: fmt.Errorf("invalid url format")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid url format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Uuid() *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			re := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
			if !re.MatchString(strInput) {
				return &Result[string]{Error: fmt.Errorf("invalid UUID")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid UUID format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) Includes(substr string) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			strSubstr := convertToString(substr)
			if !strings.Contains(strInput, strSubstr) {
				return &Result[string]{Error: fmt.Errorf("must include %s", strSubstr)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) StartsWith(prefix string) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			strPrefix := convertToString(prefix)
			if !strings.HasPrefix(strInput, strPrefix) {
				return &Result[string]{Error: fmt.Errorf("must start with %s", strPrefix)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *StringSchema[string]) EndsWith(suffix string) *Schema[string] {
	return &Schema[string]{
		Parse: func(input any) *Result[string] {
			strInput := convertToString(input)
			strSuffix := convertToString(suffix)
			if !strings.HasSuffix(strInput, strSuffix) {
				return &Result[string]{Error: fmt.Errorf("must end with %s", strSuffix)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &Result[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &Result[string]{Value: revertedStr.(string)}
		},
	}
}
