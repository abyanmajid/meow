package meow

type MeowSchema[T any] struct {
	Parse func(input any) (T, error)
	Label string
}

func (schema *MeowSchema[T]) PathLabel(label string) *MeowSchema[T] {
	schema.Label = label
	return schema
}
