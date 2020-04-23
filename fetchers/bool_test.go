package fetchers

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"errors"
)

func TestBool_Fetch(t *testing.T) {
	envKey := "TEST_VALUE"

	tcs := []struct {
		Name     string
		Default  bool
		Expected bool
		Before   func()
		ErrCheck func(t *testing.T, err error)
	}{
		{
			Name:     "happy path",
			Default:  false,
			Expected: true,
			Before: func() {
				os.Setenv(envKey, "true")
			},
		},
		{
			Name:     "not set",
			Default:  true,
			Expected: true,
			ErrCheck: func(t *testing.T, err error) {
				require.True(t, errors.Is(err, ErrNotSet))
			},
		},
		{
			Name:     "parse error",
			Default:  true,
			Expected: true,
			Before: func() {
				os.Setenv(envKey, "")
			},
			ErrCheck: func(t *testing.T, err error) {
				require.NotNil(t, err)
				require.False(t, errors.Is(err, ErrNotSet))
			},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			defer os.Unsetenv(envKey)

			if tc.Before != nil {
				tc.Before()
			}

			var v bool
			err := NewBool(&v, tc.Default).Fetch(envKey)

			if tc.ErrCheck != nil {
				tc.ErrCheck(t, err)
			} else {
				require.Nil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestBool_Value(t *testing.T) {
	v := true
	require.Nil(t, Bool{}.Value())
	require.Equal(t, v, Bool{Pointer: &v}.Value())
}
