package parsers

import (
	"strconv"
)

// Float64 is the env.Parser implementation for the float64 type.
type Float64 struct {
	Pointer *float64
	Default float64
}

// NewFloat64 returns a Float64 struct as value which holds a pointer
// and a default value.
func NewFloat64(ptr *float64, def float64) Float64 {
	return Float64{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to a float64 using strconv.ParseFloat
// and returns an error if that fails. Otherwise it sets the value
// of the pointer to the result of the conversion.
func (f Float64) Parse(str string) error {
	conv, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*f.Pointer = conv
	return err
}

// SetToDefault sets the value of the pointer to the default.
func (f Float64) SetToDefault() {
	*f.Pointer = f.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (f Float64) Value() interface{} {
	if f.Pointer == nil {
		return nil
	}

	return *f.Pointer
}
