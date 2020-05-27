package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"errors"
)

type MockParser struct {
	Err error
	Val interface{}
}

func (p MockParser) Parse(string) error {
	return p.Err
}

func (p MockParser) SetToDefault() {}

func (p MockParser) Value() interface{} {
	return p.Val
}

func TestVar_Fetch(t *testing.T) {
	testKey := t.Name()
	setEnv := func() {
		// the value could be anything as it just gets passed to the mock parser
		// really this is just to test the logic between the value being set or not
		os.Setenv(testKey, "i could be anything")
	}

	tcs := []struct {
		Name        string
		Var         Var
		Success     bool // whether or not the success logger is used
		ExpectedMsg string
		Before      func()
	}{
		{
			Name: "success",
			Var: Var{}.WithKey(testKey).WithParser(MockParser{
				Err: nil,
				Val: "test",
			}),
			Success:     true,
			ExpectedMsg: fmt.Sprintf("set %v=test", testKey),
			Before:      setEnv,
		},
		{
			Name: "success using default",
			Var: Var{}.WithKey(testKey).WithParser(MockParser{
				Err: nil,
				Val: "test",
			}),
			Success:     true,
			ExpectedMsg: fmt.Sprintf("set %v=test, default was used - %v", testKey, ErrNotSet),
		},
		{
			Name: "success with sensitive value",
			Var: Var{}.WithKey(testKey).WithParser(MockParser{
				Err: nil,
				Val: "test",
			}).MakeSensitive(),
			Success:     true,
			ExpectedMsg: fmt.Sprintf("set %v=****, value is filtered", testKey),
			Before:      setEnv,
		},
		{
			Name: "failure",
			Var: Var{}.WithKey(testKey).WithParser(MockParser{
				Err: errors.New("mock error"),
				Val: "test",
			}),
			Success:     false,
			ExpectedMsg: fmt.Sprintf("set %v=test, default was used - error: mock error", t.Name()),
			Before:      setEnv,
		},
		{
			Name: "failure using default",
			Var: Var{}.WithKey(testKey).WithParser(MockParser{
				Err: nil,
				Val: "test",
			}).LogNotSetAsFailure(),
			Success:     false,
			ExpectedMsg: fmt.Sprintf("set %v=test, default was used - %v", t.Name(), ErrNotSet),
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
			defer os.Unsetenv(testKey)

			if tc.Before != nil {
				tc.Before()
			}

			var called bool
			msgCheck := func(format string, args ...interface{}) {
				require.Equal(t, tc.ExpectedMsg, fmt.Sprintf(format, args...))
				called = true
			}

			wrongLogger := func(string, ...interface{}) {
				t.Fatalf("wrong logger called")
			}

			v := tc.Var
			if tc.Success {
				v.SuccessLogger = msgCheck
				v.FailureLogger = wrongLogger
			} else {
				v.SuccessLogger = wrongLogger
				v.FailureLogger = msgCheck
			}

			v.Fetch()

			require.True(t, called)
		})
	}
}

func TestVar_WithDefaultLogger(t *testing.T) {
	var called bool
	v := Var{}.WithDefaultLogger(func(string, ...interface{}) {
		called = true
	})

	v.WithParser(MockParser{}).Fetch()

	require.True(t, called)

	called = false

	v.WithParser(MockParser{Err: errors.New("")}).Fetch()

	require.True(t, called)
}

func TestVar_WithSuccessLogger(t *testing.T) {
	testKey := t.Name()

	tcs := map[string]struct {
		Var    Var
		Called bool
		Before func()
	}{
		"when env set and parser is successful": {
			Var:    Var{}.WithParser(MockParser{}),
			Called: true,
			Before: func() {
				os.Setenv(testKey, "any value")
			},
		},
		"when env set and parser is not successful": {
			Var:    Var{}.WithParser(MockParser{Err: errors.New("")}),
			Called: false,
			Before: func() {
				os.Setenv(testKey, "any value")
			},
		},
		"when env not set and not set is not a failure": {
			Var:    Var{}.WithParser(MockParser{}),
			Called: true,
		},
		"when env not set and not set is a failure": {
			Var:    Var{}.WithParser(MockParser{}).LogNotSetAsFailure(),
			Called: false,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			defer os.Unsetenv(testKey)

			if tc.Before != nil {
				tc.Before()
			}

			var called bool
			tc.Var.WithKey(testKey).WithSuccessLogger(func(string, ...interface{}) {
				called = true
			}).Fetch()

			require.Equal(t, tc.Called, called)
		})
	}
}

func TestVar_WithFailureLogger(t *testing.T) {
	testKey := t.Name()

	tcs := map[string]struct {
		Var    Var
		Called bool
		Before func()
	}{
		"when env set and parser is successful": {
			Var:    Var{}.WithParser(MockParser{}),
			Called: false,
			Before: func() {
				os.Setenv(testKey, "any value")
			},
		},
		"when env set and parser is not successful": {
			Var:    Var{}.WithParser(MockParser{Err: errors.New("")}),
			Called: true,
			Before: func() {
				os.Setenv(testKey, "any value")
			},
		},
		"when env not set and not set is not a failure": {
			Var:    Var{}.WithParser(MockParser{}),
			Called: false,
		},
		"when env not set and not set is a failure": {
			Var:    Var{}.WithParser(MockParser{}).LogNotSetAsFailure(),
			Called: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			defer os.Unsetenv(testKey)

			if tc.Before != nil {
				tc.Before()
			}

			var called bool
			tc.Var.WithKey(testKey).WithFailureLogger(func(string, ...interface{}) {
				called = true
			}).Fetch()

			require.Equal(t, tc.Called, called)
		})
	}
}
