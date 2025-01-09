package x_internal_test

import (
	"fmt"
	"testing"

	x_internal "github.com/abyanmajid/z/internal"
	"github.com/stretchr/testify/assert"
)

func TestAddRule(t *testing.T) {
	schema := &x_internal.Schema[int]{Path: "test123"}
	rule := func(value int) *x_internal.Result[int] {
		return schema.NewErrorResult(fmt.Sprintf("im sad %v", value))
	}

	assert.Equal(t, 0, len(schema.Rules))
	schema.AddRule(rule)
	assert.Equal(t, 1, len(schema.Rules))
}

func TestNewSuccessResult(t *testing.T) {
	schema := &x_internal.Schema[int]{Path: "im sad"}
	result := schema.NewSuccessResult()

	assert.True(t, true)
	assert.Equal(t, "im sad", result.Path)
	assert.Empty(t, result.Errors)
}

func TestNewErrorResult(t *testing.T) {
	schema := &x_internal.Schema[int]{Path: "im sad"}
	errorMessage := "an eror occurred"
	result := schema.NewErrorResult(errorMessage)

	assert.False(t, result.Success)
	assert.Equal(t, "im sad", result.Path)
	assert.Equal(t, []string{errorMessage}, result.Errors)
}
