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

func (sc *StringSchema) Parse(value interface{}) *core.Result[string] {
	valueStr, isString := value.(string)
	if !isString {
		return sc.Base.NewErrorResult("Must be a string.")
	}

	return sc.Base.ParseGeneric(valueStr)
}

func (sc *StringSchema) ParseTyped(value string) *core.Result[string] {
	return sc.Base.ParseGeneric(value)
}

func (sc *StringSchema) Min(minLength int) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) < minLength {
			errorMessage := fmt.Sprintf("Must be longer than %d characters in length", minLength)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Max(maxLength int) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) > maxLength {
			errorMessage := fmt.Sprintf("Must be shorter than %d characters in length", maxLength)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Length(length int) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) != length {
			errorMessage := fmt.Sprintf("Must be exactly %d characters long", length)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Email() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if !regexp.MustCompile(emailRegex).MatchString(value) {
			return sc.Base.NewErrorResult("Must be a valid email address")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) URL() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			return sc.Base.NewErrorResult("Must be a valid URL")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Regex(regex *regexp.Regexp) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !regex.MatchString(value) {
			return sc.Base.NewErrorResult("Must match the required pattern")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Includes(substr string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.Contains(value, substr) {
			errorMessage := fmt.Sprintf("Must include '%s'", substr)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) StartsWith(prefix string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasPrefix(value, prefix) {
			errorMessage := fmt.Sprintf("Must start with '%s'", prefix)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) EndsWith(suffix string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasSuffix(value, suffix) {
			errorMessage := fmt.Sprintf("Must end with '%s'", suffix)
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Date() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("2006-01-02", value)
		if err != nil {
			return sc.Base.NewErrorResult("Must follow a valid date format")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Time() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("15:04:05", value)
		if err != nil {
			return sc.Base.NewErrorResult("Must follow a valid time format")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) IP() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if net.ParseIP(value) == nil {
			return sc.Base.NewErrorResult("Must be a valid IP address")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) CIDR() *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, _, err := net.ParseCIDR(value)
		if err != nil {
			return sc.Base.NewErrorResult("Must be of valid CIDR notation")
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}
