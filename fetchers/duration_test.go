package fetchers

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"errors"
)

func TestDuration_Fetch(t *testing.T) {
	envKey := "TEST_VALUE"

	tcs := []struct {
		Name     string
		Default  time.Duration
		Expected time.Duration
		Before   func()
		ErrCheck func(t *testing.T, err error)
	}{
		{
			Name:     "happy path",
			Default:  3 * time.Hour,
			Expected: 4 * time.Second,
			Before: func() {
				os.Setenv(envKey, "4s")
			},
		},
		{
			Name:     "not set",
			Default:  3 * time.Hour,
			Expected: 3 * time.Hour,
			ErrCheck: func(t *testing.T, err error) {
				require.True(t, errors.Is(err, ErrNotSet))
			},
		},
		{
			Name:     "parse error",
			Default:  3 * time.Hour,
			Expected: 3 * time.Hour,
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

			var v time.Duration
			err := NewDuration(&v, tc.Default).Fetch(envKey)

			if tc.ErrCheck != nil {
				tc.ErrCheck(t, err)
			} else {
				require.Nil(t, err)
			}

			require.Equal(t, tc.Expected, v)
		})
	}
}

func TestDuration_Value(t *testing.T) {
	v := 5 * time.Second
	require.Nil(t, Duration{}.Value())
	require.Equal(t, v, Duration{Pointer: &v}.Value())
}
