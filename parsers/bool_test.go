package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBool_Parse(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected bool
		String   string
		Success  bool
	}{
		{
			Name:     "happy path",
			Expected: true,
			String:   "true",
			Success:  true,
		},
		{
			Name:     "parse error",
			Expected: false,
			String:   "",
			Success:  false,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			var v bool
			err := NewBool(&v, true).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestBool_SetToDefault(t *testing.T) {
	var v bool
	def := true
	NewBool(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestBool_Value(t *testing.T) {
	v := true
	require.Nil(t, Bool{}.Value())
	require.Equal(t, v, Bool{Pointer: &v}.Value())
}
