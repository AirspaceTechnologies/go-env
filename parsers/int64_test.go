package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected int64
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
			var v int64
			err := NewInt64(&v, 42).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestInt64_SetToDefault(t *testing.T) {
	var v int64
	def := int64(43)
	NewInt64(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestInt64_Value(t *testing.T) {
	v := int64(55)
	require.Nil(t, Int64{}.Value())
	require.Equal(t, v, Int64{Pointer: &v}.Value())
}
