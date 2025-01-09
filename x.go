package x

import x_internal "github.com/abyanmajid/z/internal"

func String(path string) *x_internal.StringSchema {
	return x_internal.NewStringSchema(path)
}
