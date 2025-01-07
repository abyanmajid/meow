package meow

import (
	"errors"
	"testing"
)

func TestMeowBaseSchema_Validate(t *testing.T) {
	tests := []struct {
		name         string
		validateFunc func(input int) error
		input        int
		wantErr      bool
	}{
		{
			name:         "No validation function set",
			validateFunc: nil,
			input:        1,
			wantErr:      false,
		},
		{
			name: "Validation function returns no error",
			validateFunc: func(input int) error {
				return nil
			},
			input:   1,
			wantErr: false,
		},
		{
			name: "Validation function returns an error",
			validateFunc: func(input int) error {
				return errors.New("validation error")
			},
			input:   1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MeowBaseSchema[int]{
				validateFunc: tt.validateFunc,
			}
			err := s.Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MeowBaseSchema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
