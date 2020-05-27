# go-env

[![GoDoc](https://godoc.org/github.com/airspacetechnologies/go-env?status.svg)](https://godoc.org/github.com/airspacetechnologies/go-env)
[![CircleCI](https://circleci.com/gh/AirspaceTechnologies/go-env.svg?style=svg)](https://circleci.com/gh/AirspaceTechnologies/go-env)
[![Test Coverage](https://api.codeclimate.com/v1/badges/5c40766e62652f91a7d1/test_coverage)](https://codeclimate.com/github/AirspaceTechnologies/go-env/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/5c40766e62652f91a7d1/maintainability)](https://codeclimate.com/github/AirspaceTechnologies/go-env/maintainability)

A simple package to fetch, parse, and log environmental variables.

## Usage

###Fetch a single environment variable:
```Go
package main

import "github.com/airspacetechnologies/go-env"

func main() {
    // fetches an environment variable called SOME_KEY and parses the string
    // into an int or sets the pointer to the default if parsing fails.
    var i int
    env.IntVar("SOME_KEY", &i, 10).Fetch()
}
```
There are many options that can be set on an `env.Var` for logging and 
configuration via setting the fields in the struct or using a chainable method.
For more information see the [docs](https://godoc.org/github.com/airspacetechnologies/go-env).

###Full Example fetching multiple variables:
```Go
package main

import (
    "log"
    "time"
    
    "github.com/airspacetechnologies/go-env"
)

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
    vars := []env.Var{
        env.BoolVar("BOOL_VALUE", &boolValue, true),
        // .MakeSensitive() will cause the value to be redacted from the logs
        env.StringVar("STRING_VALUE", &stringValue, "default").MakeSensitive(),
        // .WithFailureLogger() will override the default failure logger
        env.DurationVar("DURATION_VALUE", &durationValue, 2*time.Hour).WithFailureLogger(log.Panicf),
        // .LogNotSetAsFailure() will cause the error logger to get called if the env is not set
        env.IntVar("INT_VALUE", &intValue, 5).LogNotSetAsFailure(),
        env.Int64Var("INT64_VALUE", &int64Value, -50),
        env.Uint64Var("UINT64_VALUE", &uint64Value, 50),
        env.Float64Var("FLOAT_VALUE", &floatValue, 5.5),
    }
    
    // fetch the variables
    env.Fetch(vars)
}
```

## Running Tests
Testing is very important and this package will maintain 100% coverage.

### Using docker:

Just run `docker-compose up` to run tests

### Without docker:

If you have `golangci-lint` installed you can use `make test`. 
Otherwise, run `go test ./... -race -cover`
