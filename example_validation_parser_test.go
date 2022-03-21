package env_test

import (
	"os"
	"strings"

	"github.com/airspacetechnologies/go-env"
	"github.com/airspacetechnologies/go-env/parsers"
	"github.com/airspacetechnologies/go-env/validators"
)

// NewPercentParser is just for convenience.
func NewPercentParser(ptr *float64, def float64) env.Parser {
	return parsers.NewFloat64(ptr, def, validators.Range[float64](0, 100))
}

// ExampleValidationPercentParser shows how to wrap parsers to add
// validation logic
func Example_validationPercentParser() {
	key := "Example_validationPercentParser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: newStdoutLogger(),
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

	// Output: set Example_validationPercentParser=1, default was used - variable was not explicitly set in env
	// set Example_validationPercentParser=2, default was used - error: the following condition failed: 101 <= 100
	// set Example_validationPercentParser=10.3
}

func NewIntArrayParser(ptr *[]int, def []int) env.Parser {
	intCsvParseFunc := func(s string) ([]int, error) {
		split := strings.Split(s, ",")
		ints := make([]int, len(split))
		for i, str := range split {
			conv, err := parsers.IntParseFunc(str)
			if err != nil {
				return nil, err
			}

			ints[i] = conv
		}

		return ints, nil
	}

	greaterThanZero := validators.Greater(0, false)

	return parsers.NewGeneric(ptr, def, intCsvParseFunc, func(v []int) error {
		return greaterThanZero(len(v))
	})
}

func Example_validationLengthIntArrayParser() {
	key := "Example_validationLengthIntArrayParser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: newStdoutLogger(),
	}
	var ints []int

	// env variable not set - uses default
	v.WithParser(NewIntArrayParser(&ints, []int{1})).Fetch()

	// env variable set to invalid value (length == 0) - uses default
	os.Setenv(key, "")

	v.WithParser(NewIntArrayParser(&ints, []int{2})).Fetch()

	// env variable set to invalid value (string not int) - uses default
	os.Setenv(key, "1,2,fail")

	v.WithParser(NewIntArrayParser(&ints, []int{3})).Fetch()

	// env variable set to csv ints of length > 0
	os.Setenv(key, "1,2,3")

	v.WithParser(NewIntArrayParser(&ints, []int{4})).Fetch()

	// Output: set Example_validationLengthIntArrayParser=[1], default was used - variable was not explicitly set in env
	// set Example_validationLengthIntArrayParser=[2], default was used - error: strconv.ParseInt: parsing "": invalid syntax
	// set Example_validationLengthIntArrayParser=[3], default was used - error: strconv.ParseInt: parsing "fail": invalid syntax
	// set Example_validationLengthIntArrayParser=[1 2 3]
}

type MyEnum int

const (
	MyEnumOne MyEnum = iota
	MyEnumTwo
	MyEnumThree
)

var (
	MyEnumsSet = []MyEnum{
		MyEnumOne,
		MyEnumTwo,
		MyEnumThree,
	}
)

func NewMyEnumParser(ptr *MyEnum, def MyEnum) env.Parser {
	myEnumParseFunc := func(s string) (MyEnum, error) {
		conv, err := parsers.IntParseFunc(s)
		return MyEnum(conv), err
	}

	return parsers.NewGeneric(ptr, def, myEnumParseFunc, validators.In(MyEnumsSet...))
}

func Example_validationMyEnumParser() {
	key := "Example_validationLengthIntArrayParser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: newStdoutLogger(),
	}
	var myEnum MyEnum

	// env variable not set - uses default
	v.WithParser(NewMyEnumParser(&myEnum, MyEnumOne)).Fetch()

	// env variable set to invalid value (length == 0) - uses default
	os.Setenv(key, "3")

	v.WithParser(NewMyEnumParser(&myEnum, MyEnumTwo)).Fetch()

	// env variable set to invalid value (string not int) - uses default
	os.Setenv(key, "2")

	v.WithParser(NewMyEnumParser(&myEnum, MyEnumThree)).Fetch()

	// Output: set Example_validationLengthIntArrayParser=0, default was used - variable was not explicitly set in env
	// set Example_validationLengthIntArrayParser=1, default was used - error: 3 is not in allowed set
	// set Example_validationLengthIntArrayParser=2
}
