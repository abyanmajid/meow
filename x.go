package x

import (
	"github.com/abyanmajid/z/internal/primitives"
)

func String(path string) *primitives.StringSchema {
	return primitives.NewStringSchema(path)
}
