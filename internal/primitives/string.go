package primitives

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"

	core "github.com/abyanmajid/v/internal"
)

type StringSchema struct {
	Schema *core.Schema[string]
}

func NewStringSchema(path string) *StringSchema {
	return &StringSchema{
		Schema: &core.Schema[string]{
			Path:  path,
			Rules: []core.Rule[string]{},
		},
	}
}

func (s *StringSchema) Parse(value interface{}) *core.Result[string] {
	valueStr, isString := value.(string)
	if !isString {
		return s.Schema.NewErrorResult("Must be a string.")
	}

	return s.Schema.ParseGeneric(valueStr)
}

func (s *StringSchema) ParseTyped(value string) *core.Result[string] {
	return s.Schema.ParseGeneric(value)
}

func (s *StringSchema) Min(minLength int) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if len(value) < minLength {
			errorMessage := fmt.Sprintf("Must be longer than %d characters in length", minLength)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Max(maxLength int) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if len(value) > maxLength {
			errorMessage := fmt.Sprintf("Must be shorter than %d characters in length", maxLength)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Length(length int) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if len(value) != length {
			errorMessage := fmt.Sprintf("Must be exactly %d characters long", length)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Email() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if !regexp.MustCompile(emailRegex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid email address")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) URL() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			return s.Schema.NewErrorResult("Must be a valid URL")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Regex(regex *regexp.Regexp) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if !regex.MatchString(value) {
			return s.Schema.NewErrorResult("Must match the required pattern")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Includes(substr string) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if !strings.Contains(value, substr) {
			errorMessage := fmt.Sprintf("Must include '%s'", substr)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) StartsWith(prefix string) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if !strings.HasPrefix(value, prefix) {
			errorMessage := fmt.Sprintf("Must start with '%s'", prefix)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) EndsWith(suffix string) *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if !strings.HasSuffix(value, suffix) {
			errorMessage := fmt.Sprintf("Must end with '%s'", suffix)
			return s.Schema.NewErrorResult(errorMessage)
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Date() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("2006-01-02", value)
		if err != nil {
			return s.Schema.NewErrorResult("Must follow a valid date format")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) Time() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("15:04:05", value)
		if err != nil {
			return s.Schema.NewErrorResult("Must follow a valid time format")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) IP() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		if net.ParseIP(value) == nil {
			return s.Schema.NewErrorResult("Must be a valid IP address")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) CIDR() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		_, _, err := net.ParseCIDR(value)
		if err != nil {
			return s.Schema.NewErrorResult("Must be of valid CIDR notation")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) UUID() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		uuidRegex := `^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
		if !regexp.MustCompile(uuidRegex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid UUID")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) NanoID() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		nanoidRegex := `^[a-zA-Z0-9_-]{21}$`
		if !regexp.MustCompile(nanoidRegex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid NanoID")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) CUID() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		cuidRegex := `^c[0-9a-z]{24}$`
		if !regexp.MustCompile(cuidRegex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid CUID")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) CUID2() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		cuid2Regex := `^[a-z][a-z0-9]*$`
		if !regexp.MustCompile(cuid2Regex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid CUID2")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}

func (s *StringSchema) ULID() *StringSchema {
	s.Schema.AddRule(func(value string) *core.Result[string] {
		ulidRegex := `^[0-9A-HJKMNP-TV-Z]{26}$`
		if !regexp.MustCompile(ulidRegex).MatchString(value) {
			return s.Schema.NewErrorResult("Must be a valid ULID")
		}
		return s.Schema.NewSuccessResult()
	})
	return s
}
