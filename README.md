# V

![Tests](https://github.com/abyanmajid/v/actions/workflows/tests.yml/badge.svg) [![codecov](https://codecov.io/gh/abyanmajid/v/branch/master/graph/badge.svg?token=PkJaofBVyv)](https://codecov.io/gh/abyanmajid/v/tree/master) [![Go Report](https://goreportcard.com/badge/abyanmajid/v)](https://goreportcard.com/report/abyanmajid/v) [![MIT License](https://img.shields.io/badge/license-GPL3-blue.svg)](https://github.com/abyanmajid/v/blob/master/LICENSE)

Simple schema validation toolkit for Golang with Zod-like API.

## Quickstart

Add `v` to your Golang project:

```
go get -u github.com/abyanmajid/v@v0.6.0
```

Start composing schemas, and use them to validate your data:

```go
type Applicant struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	LinkedIn     string    `json:"linkedin"`
	University   string    `json:"university"`
	WAM          int       `json:"wam"`
	HasGraduated bool      `json:"has_graduated"`
	Courseworks  []string  `json:"courseworks"`
	Born         time.Time `json:"born"`
}

var unis = []string{"USYD", "UMELB", "UNSW"}

func isValidApplicant(a *Applicant) bool {
  // Name must be a string [1, 128] characters long.
	name := v.String("Name").Min(1).Max(128).Parse(a.Name)

  // Email must be of valid email format
	email := v.String("Email").Email().Parse(a.Email)

  // URL must be of valid URL format
	linkedIn := v.String("LinkedIn").URL().Parse(a.LinkedIn)

  // University must only be "USYD", "UMELB", "UNSW"
	university := v.Enum("University", unis).Parse(a.University)

  // WAM is an integer
  // - greater than or equal (GTE) to 0, and
  // - less than or equal to (LTE) 100
	wam := v.Integer("WAM").Gte(0).Lte(100).Parse(a.WAM)

  // Has graduated? Yes (true) or No (false)?
	hasGraduated := v.Boolean("HasGraduated").Parse(a.HasGraduated)

  // Collection of courseworks taken
	courseworks := v.Array("Courseworks", v.String("Coursework").Schema).Parse(a.Courseworks)

  // Day and time born following standard ISO format
	born := v.Date("Born").Max(time.Now()).Parse(a.Born)

  // Assert whether all validation checks has passed
	return name.Ok && email.Ok && linkedIn.Ok && university.Ok &&
		wam.Ok && hasGraduated.Ok && courseworks.Ok && born.Ok
}
```

The `Result` struct also has other fields apart from `Ok`, including:

- `Errors` (array of validation checks not passed)
- `Value` (post-validation typesafe value, which is useful if your original data is of variable type e.g., `any`/`interface{}`)
- `Path` (the name you passed in for the data, e.g., `MyPath123` in `v.String("MyPath123")`)

## API Reference

### Primitives

<details>
<summary>String</summary>

#### Strings: `String(path string)`

- Length validators: `Min(minLength int)`, `Max(maxLength int)`, `Length(length int)`
- Email validator: `Email()`
- URL validator: `URL()`
- Regex validator: `Regex(regex *regexp.Regexp)`
- Substring validators: `Includes(substr string)`, `StartsWith(prefix string)`, `EndsWith(suffix string)`
- Datetime validators: `Date()`, `Time()`
- IP validators: `IP()`, `CIDR()`
- ID validators: `UUID()`, `NanoID()`, `CUID()`, `CUID2()`, `ULID()`

</details>

<details>
<summary>Number (Integer, Float)</summary>

#### Number: `Integer(path string)`, `Float(path string)`

- Comparison validators: `Gt(lowerBound T)`, `Gte(lowerBound T)`, `Lt(upperBound T)`, `Lte(upperBund T)`,
- Sign validators: `Positive()`, `NonNegative()`, `Negative()`, `NonPositive()`,
- Multiplicity validator: `MultipleOf(step T)`
- Infinity validator: `Finite()`

Where `T` is `int` for `Integer` and `float64` for `Float`

</details>

<details>
<summary>Boolean</summary>

#### `Boolean(path string)`

No chaining validators available.

</details>

<details>
<summary>Date</summary>

#### `Date(path string)`

- Range validators: `Min(earliest time.Time)`, `Max(time.Time)`

</details>

<details>
<summary>Nil, Any, Never</summary>

#### `Nil(path string)`

No chaining validators available.

#### `Any(path string)`

No chaining validators available.

#### `Never(path string)`

No chaining validators available.

</details>

### Coercion

If you'd given a data of which type that is encoded in another type, you can coerce the data to be encoded in the type you need:

```go
now := v.Coerce.Date("Date").Parse("2025-01-10") // 2025-01-10 00:00:00 +0000 UTC
you_are_funny := v.Coerce.Boolean("You Are Funny").Parse("false") // false
nice := v.Coerce.Integer("Nice").Parse("69") // 69
nice2 := v.Coerce.String("Nice2").Parse(69) // "69"
```

### Composites

**Array:** You can define an array schema using `Array(path string, inner *core.Schema[T])`, as shown in the following example:

```go
projects := v.Array("Projects", v.String("Project").Schema)
stockPrices := v.Array("Stock Prices", v.Float("Stock Price").Schema)
```

**Struct:** This package currently doesn't provide a way to parse structs.

### Enums

You can define an enum using `Enum(path string, allowedValues []T)`, for any primitive type `T`

```go
majors := []string{"Computer Science", "Software Development", "Data Science", "Cybersecurity"}
majorsEnum := v.Enum("Majors", majors)
```

## Literals

You can define a literal using `Literal[T comparable](path string, literalValue T)` for any primitive type `T`

```go
usPresident2025 := v.Literal("US President 2025", "Donald Trump")
```

## License

This package is GPL-3.0 licensed.
