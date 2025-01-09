package primitives

import (
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

func (sc *StringSchema) Min(minLength int, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) < minLength {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Max(maxLength int, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) > maxLength {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Length(length int, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if len(value) != length {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Email(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if !regexp.MustCompile(emailRegex).MatchString(value) {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) URL(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Regex(regex *regexp.Regexp, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !regex.MatchString(value) {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Includes(substr, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.Contains(value, substr) {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) StartsWith(prefix, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasPrefix(value, prefix) {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) EndsWith(suffix, errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if !strings.HasSuffix(value, suffix) {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Date(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("2006-01-02", value)
		if err != nil {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) Time(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, err := time.Parse("15:04:05", value)
		if err != nil {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) IP(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		if net.ParseIP(value) == nil {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}

func (sc *StringSchema) CIDR(errorMessage string) *StringSchema {
	sc.Base.AddRule(func(value string) *core.Result[string] {
		_, _, err := net.ParseCIDR(value)
		if err != nil {
			return sc.Base.NewErrorResult(errorMessage)
		}
		return sc.Base.NewSuccessResult()
	})
	return sc
}
