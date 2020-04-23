package fetchers

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"errors"
)

func TestString_Fetch(t *testing.T) {
	envKey := "TEST_VALUE"

	tcs := []struct {
		Name     string
		Default  string
		Expected string
		Before   func()
		ErrCheck func(t *testing.T, err error)
	}{
		{
			Name:     "happy path",
			Default:  "5",
			Expected: "10",
			Before: func() {
				os.Setenv(envKey, "10")
			},
		},
		{
			Name:     "not set",
			Default:  "5",
			Expected: "5",
			ErrCheck: func(t *testing.T, err error) {
				require.True(t, errors.Is(err, ErrNotSet))
			},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			defer os.Unsetenv(envKey)

			if tc.Before != nil {
				tc.Before()
			}

			var v string
			err := NewString(&v, tc.Default).Fetch(envKey)

			if tc.ErrCheck != nil {
				tc.ErrCheck(t, err)
			} else {
				require.Nil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestString_Value(t *testing.T) {
	v := "test"
	require.Nil(t, String{}.Value())
	require.Equal(t, v, String{Pointer: &v}.Value())
}
