# go-env

A simple package to fetch environmental variables

## Usage

Example:
```Go
import "github.com/airspacetechnologies/go-tools/env"

var (
    boolValue     bool
    stringValue   string
    durationValue time.Duration
    intValue      int
    int64Value    int64
    uint64Value   uint64
    floatValue    float64
)

func init() {
    // set log functions
    env.DefaultLogger = log.Printf
    env.FailureLogger = log.Fatalf

    // define vars
    vars := []Var{
        BoolVar("BOOL_VALUE", &boolValue, true),
        // .MakeSensitive() will cause the value to be redacted from the logs
        StringVar("STRING_VALUE", &stringValue, "default").MakeSensitive(),
        // .WithFailureLogger() will override the default failure logger
        DurationVar("DURATION_VALUE", &durationValue, 2*time.Hour).WithFailureLogger(log.Panicf),
        // .LogNotSetAsFailure() will cause the error logger to get called if the env is not set
        IntVar("INT_VALUE", &intValue, 5).LogNotSetAsFailure(),
        Int64Var("INT64_VALUE", &int64Value, -50),
        Uint64Var("UINT64_VALUE", &uint64Value, 50),
        Float64Var("FLOAT_VALUE", &floatValue, 5.5),
    }
    
    // fetch the variables
    Fetch(vars)
}
```

## Running Tests

Just run `docker-compose up` to run tests
