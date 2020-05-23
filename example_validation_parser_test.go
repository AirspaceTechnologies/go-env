package env_test

import (
	"errors"
	"fmt"
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

	var i float64

	// env variable not set
	env.Var{Key: key, Parser: NewPercentParser(&i, 1)}.Fetch()
	fmt.Println(i) // default value used (1), Parse not called

	// env variable set to negative (invalid) string
	os.Setenv(key, "-1")

	i = 0
	env.Var{Key: key, Parser: NewPercentParser(&i, 2)}.Fetch()
	fmt.Println(i) // default value used (2), Parse failed

	// env variable set to a valid hex string
	os.Setenv(key, "10.3")

	i = 0
	env.Var{Key: key, Parser: NewPercentParser(&i, 3)}.Fetch()
	fmt.Println(i) // parsed successfully (10.3)

	// Output: 1
	// 2
	// 10.3
}
