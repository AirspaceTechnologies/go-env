package env_test

import (
	"errors"
	"log"
	"os"

	"github.com/airspacetechnologies/go-env"
	"github.com/airspacetechnologies/go-env/parsers"
)

// PercentParser wraps the Float64 parser to add validation.
type PercentParser struct {
	parsers.Float64
}

// NewPercentParser is just for convenience.
func NewPercentParser(ptr *float64, def float64) PercentParser {
	return PercentParser{
		Float64: parsers.NewFloat64(ptr, def),
	}
}

// Parse calls the wrapped parser's Parse and then validates upon success.
func (p PercentParser) Parse(str string) error {
	if err := p.Float64.Parse(str); err != nil {
		return err
	}

	// get value
	v := *p.Pointer

	// do validations
	if v < 0 || v > 100 {
		return errors.New("percent is out of bounds")
	}

	return nil
}

// ExampleValidationParser shows how to wrap parsers to add
// validation logic
func Example_validationParser() {
	key := "Example_parserValidation"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: log.New(os.Stdout, "", 0).Printf, // sends logs to os.Stdout for examples
	}
	var i float64

	// env variable not set - uses default
	v.WithParser(NewPercentParser(&i, 1)).Fetch()

	// env variable set to invalid value (> 100) - uses default
	os.Setenv(key, "101")

	v.WithParser(NewPercentParser(&i, 2)).Fetch()

	// env variable set to a number between 0 and 100
	os.Setenv(key, "10.3")

	v.WithParser(NewPercentParser(&i, 3)).Fetch()

	// Output: set Example_parserValidation=1, default was used - variable was not explicitly set in env
	// set Example_parserValidation=2, default was used - error: percent is out of bounds
	// set Example_parserValidation=10.3
}
