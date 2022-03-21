package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected int
		String   string
		Success  bool
	}{
		{
			Name:     "happy path",
			Expected: 10,
			String:   "10",
			Success:  true,
		},
		{
			Name:     "parse error",
			Expected: 0,
			String:   "",
			Success:  false,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			var v int
			err := NewInt(&v, 402).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestInt_SetToDefault(t *testing.T) {
	var v int
	def := 402
	NewInt(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestInt_Value(t *testing.T) {
	v := 55
	require.Nil(t, NewInt(nil, 0).Value())
	require.Equal(t, v, NewInt(&v, 0).Value())
}
