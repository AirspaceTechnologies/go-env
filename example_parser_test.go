package env_test

import (
	"encoding/hex"
	"log"
	"os"

	"github.com/airspacetechnologies/go-env"
)

// HexParser is an example of how to make a custom parser.
type HexParser struct {
	Pointer *[]byte
	Default []byte
}

// NewHexParser is not required, it is just for convenience.
func NewHexParser(ptr *[]byte, def []byte) HexParser {
	return HexParser{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string and sets the pointer upon success.
// If it fails it returns an error.
func (p HexParser) Parse(str string) error {
	// convert string
	conv, err := hex.DecodeString(str)
	if err != nil {
		// return error if conversion failed
		return err
	}

	// set the value of the ppinter
	*p.Pointer = conv
	return nil
}

// SetToDefault gets called if the environmental variable was
// not set or if Parse returned an error. It just sets the
// value of the pointer.
func (p HexParser) SetToDefault() {
	*p.Pointer = p.Default
}

// Value returns the value of the pointer or nil as an
// interface{} and is used for logging.
func (p HexParser) Value() interface{} {
	if p.Pointer == nil {
		return nil
	}

	return *p.Pointer
}

// ExampleParser shows how to make a custom parser.
func Example_parser() {
	key := "Example_parser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: log.New(os.Stdout, "", 0).Printf, // sends logs to os.Stdout for examples
	}
	var b []byte

	// env variable not set - uses default
	v.WithParser(NewHexParser(&b, []byte{1})).Fetch()

	// env variable set to invalid hex string - uses default
	os.Setenv(key, "gg")

	v.WithParser(NewHexParser(&b, []byte{2})).Fetch()

	// env variable set to a valid hex string
	os.Setenv(key, "ff")

	v.WithParser(NewHexParser(&b, []byte{3})).Fetch()

	// Output: set Example_parser=[1], default was used - variable was not explicitly set in env
	// set Example_parser=[2], default was used - error: encoding/hex: invalid byte: U+0067 'g'
	// set Example_parser=[255]
}
