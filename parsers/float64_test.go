package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloat64_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected float64
		String   string
		Success  bool
	}{
		{
			Name:     "happy path",
			Expected: 10.7,
			String:   "10.7",
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
			var v float64
			err := NewFloat64(&v, 30.2).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestFloat64_SetToDefault(t *testing.T) {
	var v float64
	def := 2.4
	NewFloat64(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestFloat64_Value(t *testing.T) {
	v := 55.5
	require.Nil(t, NewFloat64(nil, 0).Value())
	require.Equal(t, v, NewFloat64(&v, 0).Value())
}
