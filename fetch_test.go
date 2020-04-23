package env

import (
	"os"
	"testing"
	"time"

	"github.com/airspacetechnologies/go-env/fetchers"

	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	defer func(logFunc LogFunc) { DefaultLogger = logFunc }(DefaultLogger)
	DefaultLogger = ZeroLogger

	var (
		boolValue     bool
		stringValue   string
		durationValue time.Duration
		intValue      int
		int64Value    int64
		uint64Value   uint64
		floatValue    float64
	)

	vars := []Var{
		BoolVar("BOOL_VALUE", &boolValue, true),
		StringVar("STRING_VALUE", &stringValue, "default"),
		DurationVar("DURATION_VALUE", &durationValue, 2*time.Hour),
		IntVar("INT_VALUE", &intValue, 5),
		Int64Var("INT64_VALUE", &int64Value, -50),
		Uint64Var("UINT64_VALUE", &uint64Value, 50),
		Float64Var("FLOAT_VALUE", &floatValue, 5.5),
	}

	// when values not set (defaults)
	Fetch(vars)

	require.True(t, boolValue)
	require.Equal(t, "default", stringValue)
	require.Equal(t, 2*time.Hour, durationValue)
	require.Equal(t, 5, intValue)
	require.Equal(t, int64(-50), int64Value)
	require.Equal(t, uint64(50), uint64Value)
	require.Equal(t, 5.5, floatValue)

	// when values are set
	vals := map[string]string{
		"BOOL_VALUE":     "false",
		"STRING_VALUE":   "test",
		"DURATION_VALUE": "4h",
		"INT_VALUE":      "10",
		"INT64_VALUE":    "-125",
		"UINT64_VALUE":   "125",
		"FLOAT_VALUE":    "10.5",
	}

	// capture current env vars to reset env after tests
	type valHolder struct {
		v  string
		ok bool
	}

	initialValues := map[string]valHolder{}
	for k := range vals {
		v, ok := os.LookupEnv(k)
		initialValues[k] = valHolder{v: v, ok: ok}
	}

	// reset env after tests
	defer func(map[string]valHolder) {
		for k, h := range initialValues {
			if h.ok {
				require.Nil(t, os.Setenv(k, h.v))
			} else {
				require.Nil(t, os.Unsetenv(k))
			}
		}
	}(initialValues)

	// set env vars
	for k, v := range vals {
		require.Nil(t, os.Setenv(k, v))
	}

	Fetch(vars)

	require.False(t, boolValue)
	require.Equal(t, "test", stringValue)
	require.Equal(t, 4*time.Hour, durationValue)
	require.Equal(t, 10, intValue)
	require.Equal(t, int64(-125), int64Value)
	require.Equal(t, uint64(125), uint64Value)
	require.Equal(t, 10.5, floatValue)
}

func TestFetchMap(t *testing.T) {
	var (
		boolValue     bool
		stringValue   string
		durationValue time.Duration
		intValue      int
		floatValue    float64
	)

	vars := map[string]Var{
		"BOOL_VALUE": {
			Fetcher: fetchers.NewBool(&boolValue, true),
		},
		"STRING_VALUE": {
			Fetcher: fetchers.NewString(&stringValue, "default"),
		},
		"DURATION_VALUE": {
			Fetcher: fetchers.NewDuration(&durationValue, 2*time.Hour),
		},
		"INT_VALUE": {
			Fetcher: fetchers.NewInt(&intValue, 5),
		},
		"FLOAT_VALUE": {
			Fetcher: fetchers.NewFloat64(&floatValue, 5.5),
		},
	}

	// when values not set (defaults)
	FetchMap(vars)

	require.True(t, boolValue)
	require.Equal(t, "default", stringValue)
	require.Equal(t, 2*time.Hour, durationValue)
	require.Equal(t, 5, intValue)
	require.Equal(t, 5.5, floatValue)

	// when values are set
	vals := map[string]string{
		"BOOL_VALUE":     "false",
		"STRING_VALUE":   "test",
		"DURATION_VALUE": "4h",
		"INT_VALUE":      "10",
		"FLOAT_VALUE":    "10.5",
	}

	// capture current env vars to reset env after tests
	type valHolder struct {
		v  string
		ok bool
	}

	initialValues := map[string]valHolder{}
	for k := range vals {
		v, ok := os.LookupEnv(k)
		initialValues[k] = valHolder{v: v, ok: ok}
	}

	// reset env after tests
	defer func(map[string]valHolder) {
		for k, h := range initialValues {
			if h.ok {
				require.Nil(t, os.Setenv(k, h.v))
			} else {
				require.Nil(t, os.Unsetenv(k))
			}
		}
	}(initialValues)

	// set env vars
	for k, v := range vals {
		require.Nil(t, os.Setenv(k, v))
	}

	FetchMap(vars)

	require.False(t, boolValue)
	require.Equal(t, "test", stringValue)
	require.Equal(t, 4*time.Hour, durationValue)
	require.Equal(t, 10, intValue)
	require.Equal(t, 10.5, floatValue)
}
