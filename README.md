# V

![Tests](https://github.com/abyanmajid/v/actions/workflows/tests.yml/badge.svg) [![codecov](https://codecov.io/gh/abyanmajid/v/branch/master/graph/badge.svg?token=PkJaofBVyv)](https://codecov.io/gh/abyanmajid/v/tree/master) [![Go Report](https://goreportcard.com/badge/abyanmajid/v)](https://goreportcard.com/report/abyanmajid/v) [![MIT License](https://img.shields.io/badge/license-GPL3-blue.svg)](https://github.com/abyanmajid/v/blob/master/LICENSE)

Simple schema validation toolkit for Golang with Zod-like API.

## Usage

Add `v` to your Golang project:

```
go get -u github.com/abyanmajid/v@v0.5.0
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
	name := v.String("Name").Min(1).Max(128).Parse(a.Name)
	email := v.String("Email").Email().Parse(a.Email)
	linkedIn := v.String("LinkedIn").URL().Parse(a.LinkedIn)
	university := v.Enum("University", unis).Parse(a.University)
	wam := v.Integer("WAM").Gte(0).Lte(100).Parse(a.WAM)
	hasGraduated := v.Boolean("HasGraduated").Parse(a.HasGraduated)
	courseworks := v.Array("Courseworks", v.String("Coursework").Schema).Parse(a.Courseworks)
	born := v.Date("Born").Max(time.Now()).Parse(a.Born)

	return name.Success && email.Success &&
		linkedIn.Success && university.Success &&
		wam.Success && hasGraduated.Success &&
		courseworks.Success && born.Success
}
```

## License

This package is GPL-3.0 licensed.
