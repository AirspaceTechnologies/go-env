package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected string
		String   string
	}{
		{
			Name:     "happy path (the only path)",
			Expected: "10",
			String:   "10",
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			var v string
			err := NewString(&v, "abc").Parse(tc.String)
			require.Nil(t, err)
			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestString_SetToDefault(t *testing.T) {
	var v string
	def := "abc"
	NewString(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestString_Value(t *testing.T) {
	v := "test"
	require.Nil(t, String{}.Value())
	require.Equal(t, v, String{Pointer: &v}.Value())
}
