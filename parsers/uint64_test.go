package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint64_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected uint64
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
			var v uint64
			err := NewUint64(&v, 43).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestUint64_SetToDefault(t *testing.T) {
	var v uint64
	def := uint64(402)
	NewUint64(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestUint64_Value(t *testing.T) {
	v := uint64(55)
	require.Nil(t, Uint64{}.Value())
	require.Equal(t, v, Uint64{Pointer: &v}.Value())
}
