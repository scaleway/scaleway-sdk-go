package injector

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestInject(t *testing.T) {

	type S struct {
		LastName  string `scw:"lastName" json:"lastName,omitempty"`
		FirstName string `scw:"firstName"`
		Age       int    `scw:"age"`
	}

	testCases := []struct {
		name     string
		target   *S
		tagName  string
		data     map[string]interface{}
		expected *S
		error    string
	}{
		{
			name:    "simple assign",
			target:  &S{},
			tagName: "scw",
			data: map[string]interface{}{
				"lastName":  "sponge",
				"firstName": "bob",
				"age":       42,
			},
			expected: &S{
				LastName:  "sponge",
				FirstName: "bob",
				Age:       42,
			},
		},
		{
			name:    "non existing tag",
			target:  &S{},
			tagName: "scw",
			data: map[string]interface{}{
				"non-existing": 42,
			},
			expected: &S{},
		},
		{
			name:    "invalid type",
			target:  &S{},
			tagName: "scw",
			data: map[string]interface{}{
				"age": "this should be an int",
			},
			expected: &S{},
			error:    "trying to assign incompatible type string to int for tag age",
		},
		{
			name:    "json style tags",
			target:  &S{},
			tagName: "json",
			data: map[string]interface{}{
				"lastName": "sponge",
			},
			expected: &S{
				LastName: "sponge",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Inject(tc.target, tc.tagName, tc.data)
			if tc.error == "" {
				testhelpers.Ok(t, err)
				testhelpers.Equals(t, tc.expected, tc.target)
			} else {
				testhelpers.Equals(t, tc.error, err.Error())
			}
		})
	}
}
