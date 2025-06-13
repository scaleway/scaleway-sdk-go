package generic

import "testing"

func TestHasField(t *testing.T) {
	tests := []struct {
		name      string
		i         any
		fieldName string
		want      bool
	}{
		{
			"native type",
			struct {
				Zone string
			}{},
			"Zone",
			true,
		},
		{
			"ptr type",
			struct {
				Zone *string
			}{},
			"Zone",
			true,
		},
		{
			"invalid case",
			struct {
				Zone *string
			}{},
			"zone",
			false,
		},
		{
			"slice",
			[]struct {
				Zone string
			}{},
			"Zone",
			true,
		},
		{
			"slice invalid case",
			[]struct {
				Zone string
			}{},
			"zone",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasField(tt.i, tt.fieldName); got != tt.want {
				t.Errorf("HasField() = %v, want %v", got, tt.want)
			}
		})
	}
}
