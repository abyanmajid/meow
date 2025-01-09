package primitives

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"

	core "github.com/abyanmajid/z/internal"
)

type StringSchema struct {
	Base *core.Schema[string]
}

func NewStringSchema(path string) *StringSchema {
	return &StringSchema{
		Base: &core.Schema[string]{
			Path:  path,
			Rules: []core.Rule[string]{},
		},
	}
}

func (s *StringSchema) Parse(value interface{}) *core.Result[string] {
	valueStr, isString := value.(string)
	if !isString {
		return s.Base.NewErrorResult("Must be a string.")
	}

	return s.Base.ParseGeneric(valueStr)
}

func (s *StringSchema) ParseTyped(value string) *core.Result[string] {
	return s.Base.ParseGeneric(value)
}

func (s *StringSchema) Min(minLength int) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) < minLength {
			errorMessage := fmt.Sprintf("Must be longer than %d characters in length", minLength)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Max(maxLength int) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) > maxLength {
			errorMessage := fmt.Sprintf("Must be shorter than %d characters in length", maxLength)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Length(length int) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) != length {
			errorMessage := fmt.Sprintf("Must be exactly %d characters long", length)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Email() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if !regexp.MustCompile(emailRegex).MatchString(value) {
			return s.Base.NewErrorResult("Must be a valid email address")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) URL() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			return s.Base.NewErrorResult("Must be a valid URL")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Regex(regex *regexp.Regexp) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if !regex.MatchString(value) {
			return s.Base.NewErrorResult("Must match the required pattern")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Includes(substr string) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.Contains(value, substr) {
			errorMessage := fmt.Sprintf("Must include '%s'", substr)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) StartsWith(prefix string) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasPrefix(value, prefix) {
			errorMessage := fmt.Sprintf("Must start with '%s'", prefix)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) EndsWith(suffix string) *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasSuffix(value, suffix) {
			errorMessage := fmt.Sprintf("Must end with '%s'", suffix)
			return s.Base.NewErrorResult(errorMessage)
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Date() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("2006-01-02", value)
		if err != nil {
			return s.Base.NewErrorResult("Must follow a valid date format")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Time() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("15:04:05", value)
		if err != nil {
			return s.Base.NewErrorResult("Must follow a valid time format")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) IP() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		if net.ParseIP(value) == nil {
			return s.Base.NewErrorResult("Must be a valid IP address")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) CIDR() *StringSchema {
	s.Base.AddRule(func(value string) *core.Result[string] {
		_, _, err := net.ParseCIDR(value)
		if err != nil {
			return s.Base.NewErrorResult("Must be of valid CIDR notation")
		}
		return s.Base.NewSuccessResult()
	})
	return s
}
