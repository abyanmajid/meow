package meow

import (
	"testing"
)

func TestPathLabel(t *testing.T) {
	schema := &MeowSchema[string]{}
	label := "testLabel"

	result := schema.PathLabel(label)

	if result.Label != label {
		t.Errorf("expected label %s, got %s", label, result.Label)
	}

	if result != schema {
		t.Errorf("expected schema to be returned")
	}
}
