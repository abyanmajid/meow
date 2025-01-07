package meow

import (
	"fmt"
	"net"
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
