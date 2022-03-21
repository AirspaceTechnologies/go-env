package parsers

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDuration_Fetch(t *testing.T) {
	tcs := []struct {
		Name     string
		Expected time.Duration
		String   string
		Success  bool
	}{
		{
			Name:     "happy path",
			Expected: 4 * time.Second,
			String:   "4s",
			Success:  true,
		},
		{
			Name:     "parse error",
			Expected: time.Duration(0),
			String:   "",
			Success:  false,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			var v time.Duration
			err := NewDuration(&v, time.Duration(100)).Parse(tc.String)

			if tc.Success {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestDuration_SetToDefault(t *testing.T) {
	var v time.Duration
	def := 3 * time.Hour
	NewDuration(&v, def).SetToDefault()
	require.Equal(t, def, v)
}

func TestDuration_Value(t *testing.T) {
	v := 5 * time.Second
	require.Nil(t, NewDuration(nil, 0).Value())
	require.Equal(t, v, NewDuration(&v, 0).Value())
}
