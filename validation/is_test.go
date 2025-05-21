package validation

import (
	"testing"
)

func Test_IsUUID(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect bool
	}{
		{"Valid UUID", "123e4567-e89b-12d3-a456-426614174000", true},
		{"Invalid UUID - wrong length", "123e4567-e89b-12d3-a456-42661417400", false},
		{"Invalid UUID - no dashes", "123e4567e89b12d3a456426614174000", false},
		{"Completely Invalid", "not-a-uuid", false},
		{"Valid UUID with zone prefix", "fr-par-1/123e4567-e89b-12d3-a456-426614174000", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsUUID(tt.input)
			if got != tt.expect {
				t.Errorf("IsUUID(%s) = %v; want %v", tt.input, got, tt.expect)
			}
		})
	}
}
