// parsers is a package containing type specific logic used by the env package via the
// env.Parser interface. Parsers are responsible for converting a string to the correct
// type, setting the value of its pointer for both the parsed value and default value,
// and is able to return the value of its pointer as an interface or nil if the pointer
// is nil.
//
// You can use the Generic[T] parser and pass it a custom ParseFunc and validator.Func
// in order to make parsers for custom types. Refer to any of the premade implentations
// or example tests in the env package to see how to do this.
//
// You can also add validations to the parser. See the ValidationParser
// example in the env package for how to easily do that.
package parsers

import (
	"github.com/airspacetechnologies/go-env/validators"
	"time"
)

// NewBool returns a bool Generic which uses strconv.ParseBool
func NewBool(ptr *bool, def bool, vfs ...validators.Func[bool]) Generic[bool] {
	return NewGeneric(ptr, def, BoolParseFunc, vfs...)
}

// NewDuration returns a duration Generic which uses time.ParseDuration
func NewDuration(ptr *time.Duration, def time.Duration, vfs ...validators.Func[time.Duration]) Generic[time.Duration] {
	return NewGeneric(ptr, def, DurationParseFunc, vfs...)
}

// NewFloat64 returns a float64 Generic which uses strconv.ParseFloat
func NewFloat64(ptr *float64, def float64, vfs ...validators.Func[float64]) Generic[float64] {
	return NewGeneric(ptr, def, Float64ParseFunc, vfs...)
}

// NewInt returns an int Generic which uses Int64ParseFunc and converts to an int
func NewInt(ptr *int, def int, vfs ...validators.Func[int]) Generic[int] {
	return NewGeneric(ptr, def, IntParseFunc, vfs...)
}

// NewInt64 returns an int64 Generic which uses strconv.ParseInt
func NewInt64(ptr *int64, def int64, vfs ...validators.Func[int64]) Generic[int64] {
	return NewGeneric(ptr, def, Int64ParseFunc, vfs...)
}

// NewString returns a string Generic
func NewString(ptr *string, def string, vfs ...validators.Func[string]) Generic[string] {
	return NewGeneric(ptr, def, StringParseFunc, vfs...)
}

// NewUint64 returns an uint64 Generic which uses strconv.ParseUint
func NewUint64(ptr *uint64, def uint64, vfs ...validators.Func[uint64]) Generic[uint64] {
	return NewGeneric(ptr, def, Uint64ParseFunc, vfs...)
}
