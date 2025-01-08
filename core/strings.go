package meow

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

func (s *MeowSchema[string]) Datetime(offset bool, precision int) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			layout := "2006-01-02T15:04:05Z07:00" // ISO 8601 format with timezone
			if precision > 0 {
				layout = "2006-01-02T15:04:05." + strings.Repeat("0", precision) + "Z07:00"
			}
			_, err := time.Parse(layout, strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid datetime format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid datetime format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Date() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			_, err := time.Parse("2006-01-02", strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid date format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid date format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Time() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			layout := "15:04:05"
			_, err := time.Parse(layout, strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid time format: %s", err)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid time format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Ip() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			if net.ParseIP(strInput) == nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid IP format: %s", strInput)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid IP format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Cidr() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			_, _, err := net.ParseCIDR(strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid CIDR format: %s", strInput)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid CIDR format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) MinLength(min int) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			if len(strInput) < min {
				return &MeowResult[string]{Error: fmt.Errorf("must be %d or more characters long", min)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) MaxLength(max int) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			if len(strInput) > max {
				return &MeowResult[string]{Error: fmt.Errorf("must be %d or fewer characters long", max)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) ExactLength(exact int) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			if len(strInput) != exact {
				return &MeowResult[string]{Error: fmt.Errorf("must be exactly %d characters long", exact)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Email() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)

			re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
			if !re.MatchString(strInput) {
				return &MeowResult[string]{Error: fmt.Errorf("invalid email address")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid email format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Url() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)

			re := `^(http|https):\/\/[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
			match, err := regexp.MatchString(re, strInput)
			if err != nil || !match {
				return &MeowResult[string]{Error: fmt.Errorf("invalid url format")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid url format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Uuid() *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			re := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
			if !re.MatchString(strInput) {
				return &MeowResult[string]{Error: fmt.Errorf("invalid UUID")}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid UUID format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) Includes(substr string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			strSubstr := convertToString(substr)
			if !strings.Contains(strInput, strSubstr) {
				return &MeowResult[string]{Error: fmt.Errorf("must include %s", strSubstr)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) StartsWith(prefix string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			strPrefix := convertToString(prefix)
			if !strings.HasPrefix(strInput, strPrefix) {
				return &MeowResult[string]{Error: fmt.Errorf("must start with %s", strPrefix)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}

func (s *MeowSchema[string]) EndsWith(suffix string) *MeowSchema[string] {
	return &MeowSchema[string]{
		Parse: func(input any) *MeowResult[string] {
			strInput := convertToString(input)
			strSuffix := convertToString(suffix)
			if !strings.HasSuffix(strInput, strSuffix) {
				return &MeowResult[string]{Error: fmt.Errorf("must end with %s", strSuffix)}
			}

			revertedStr, err := convertFromString[string](strInput)
			if err != nil {
				return &MeowResult[string]{Error: fmt.Errorf("invalid input format: %s", err)}
			}

			return &MeowResult[string]{Value: revertedStr.(string)}
		},
	}
}
