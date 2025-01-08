package meow

type Result[T any] struct {
	Path  string
	Error error
	Value T
}

type Schema[T any] struct {
	Optional bool
	Parse    func(input any) *Result[T]
}

type StructSchema[T any] struct {
	Rules *Schema[T]
}

type StringSchema[T any] struct {
	Rules *Schema[T]
}

type NumberSchema[T any] struct {
	Rules *Schema[T]
}

type BooleanSchema[T any] struct {
	Rules *Schema[T]
}

type DateSchema[T any] struct {
	Rules *Schema[T]
}

type NilSchema[T any] struct {
	Rules *Schema[T]
}

type AnySchema[T any] struct {
	Rules *Schema[T]
}

type NeverSchema[T any] struct {
	Rules *Schema[T]
}
