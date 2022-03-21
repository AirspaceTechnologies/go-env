package env_test

import (
	"log"
	"time"

	"github.com/airspacetechnologies/go-env"
)

func ExampleFetchWithConfig() {
	var (
		boolValue     bool
		stringValue   string
		durationValue time.Duration
		intValue      int
		int64Value    int64
		uint64Value   uint64
		floatValue    float64
	)

	// define vars
	vars := []env.Var{
		env.BoolVar("BOOL_VALUE", &boolValue, true),
		// .MakeSensitive() will cause the value to be redacted from the logs
		env.StringVar("STRING_VALUE", &stringValue, "default").MakeSensitive(),
		// .WithFailureLogger() will override the default failure logger
		env.DurationVar("DURATION_VALUE", &durationValue, 2*time.Hour).WithFailureLogger(log.Fatalf),
		// .LogNotSetAsFailure() will cause the error logger to get called if the env is not set
		env.IntVar("INT_VALUE", &intValue, 5).LogNotSetAsFailure(),
		env.Int64Var("INT64_VALUE", &int64Value, -50),
		env.Uint64Var("UINT64_VALUE", &uint64Value, 50),
		env.Float64Var("FLOAT_VALUE", &floatValue, 5.5),
	}

	// set log function to log to stdout for examples
	config := env.Config{
		DefaultLogger: newStdoutLogger(),
	}

	// fetch the variables
	env.FetchWithConfig(vars, config)

	// Output: set BOOL_VALUE=true, default was used - variable was not explicitly set in env
	// set STRING_VALUE=****, value is filtered, default was used - variable was not explicitly set in env
	// set DURATION_VALUE=2h0m0s, default was used - variable was not explicitly set in env
	// set INT_VALUE=5, default was used - variable was not explicitly set in env
	// set INT64_VALUE=-50, default was used - variable was not explicitly set in env
	// set UINT64_VALUE=50, default was used - variable was not explicitly set in env
	// set FLOAT_VALUE=5.5, default was used - variable was not explicitly set in env
}
