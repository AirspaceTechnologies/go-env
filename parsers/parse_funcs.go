package parsers

import (
	"strconv"
	"time"
)

var (
	// BoolParseFunc uses strconv.ParseBool
	BoolParseFunc = strconv.ParseBool

	// DurationParseFunc uses time.ParseDuration
	DurationParseFunc = time.ParseDuration

	// Float64ParseFunc uses strconv.ParseFloat with bit size of 64
	Float64ParseFunc = func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}

	// Int64ParseFunc uses strconv.ParseInt with base 10 and bit size of 64
	Int64ParseFunc = func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	}

	// IntParseFunc uses Int64ParseFunc and converts to int
	IntParseFunc = func(s string) (int, error) {
		i64, err := Int64ParseFunc(s)
		return int(i64), err
	}

	// StringParseFunc just passes through the string with no error
	StringParseFunc = func(s string) (string, error) {
		return s, nil
	}

	// Uint64ParseFunc uses strconv.ParseUint with base 10 and bit size of 64
	Uint64ParseFunc = func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	}
)

// ParseFunc is any function that takes in a string and outputs ant type along with an error
type ParseFunc[T any] func(string) (T, error)
