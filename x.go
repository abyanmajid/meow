package x

import core "github.com/abyanmajid/z/internal"

func String(path string) *core.StringSchema {
	return core.NewStringSchema(path)
}
