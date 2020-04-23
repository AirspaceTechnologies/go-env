package env

import (
	"fmt"
	"testing"

	"github.com/airspacetechnologies/go-env/fetchers"

	"github.com/stretchr/testify/require"

	"errors"
)

type MockFetcher struct {
	Err error
	Val interface{}
}

func (f MockFetcher) Fetch(string) error {
	return f.Err
}

func (f MockFetcher) Value() interface{} {
	return f.Val
}

func TestVar_Fetch(t *testing.T) {
	testKey := "TEST_KEY"

	tcs := []struct {
		Name        string
		Var         Var
		Success     bool
		ExpectedMsg string
	}{
		{
			Name: "success",
			Var: Var{}.WithKey(testKey).WithFetcher(MockFetcher{
				Err: nil,
				Val: "test",
			}),
			Success:     true,
			ExpectedMsg: "set TEST_KEY=test",
		},
		{
			Name: "success using default",
			Var: Var{}.WithKey(testKey).WithFetcher(MockFetcher{
				Err: fetchers.ErrNotSet,
				Val: "test",
			}),
			Success:     true,
			ExpectedMsg: fmt.Sprintf("set TEST_KEY=test, default was used - %v", fetchers.ErrNotSet),
		},
		{
			Name: "success with sensitive value",
			Var: Var{}.WithKey(testKey).WithFetcher(MockFetcher{
				Err: nil,
				Val: "test",
			}).MakeSensitive(),
			Success:     true,
			ExpectedMsg: "set TEST_KEY=****, value is filtered",
		},
		{
			Name: "failure",
			Var: Var{}.WithKey(testKey).WithFetcher(MockFetcher{
				Err: errors.New("mock error"),
				Val: "test",
			}),
			Success:     false,
			ExpectedMsg: "set TEST_KEY=test, err: mock error",
		},
		{
			Name: "failure using default",
			Var: Var{}.WithKey(testKey).WithFetcher(MockFetcher{
				Err: fetchers.ErrNotSet,
				Val: "test",
			}).LogNotSetAsFailure(),
			Success:     false,
			ExpectedMsg: fmt.Sprintf("set TEST_KEY=test, default was used - %v", fetchers.ErrNotSet),
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v %v", i, tc.Name), func(t *testing.T) {
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

	v.WithFetcher(MockFetcher{}).Fetch()

	require.True(t, called)

	called = false

	v.WithFetcher(MockFetcher{Err: errors.New("")}).Fetch()

	require.True(t, called)
}

func TestVar_WithSuccessLogger(t *testing.T) {
	var called bool
	v := Var{}.WithSuccessLogger(func(string, ...interface{}) {
		called = true
	})

	v.WithFetcher(MockFetcher{Err: errors.New("")}).Fetch()

	require.False(t, called)

	v.WithFetcher(MockFetcher{}).Fetch()

	require.True(t, called)
}

func TestVar_WithFailureLogger(t *testing.T) {
	var called bool
	v := Var{}.WithFailureLogger(func(string, ...interface{}) {
		called = true
	})

	v.WithFetcher(MockFetcher{}).Fetch()

	require.False(t, called)

	v.WithFetcher(MockFetcher{Err: errors.New("")}).Fetch()

	require.True(t, called)
}
