package parsers

import (
	"strconv"
)

// Bool is the env.Parser implementation for the bool type.
type Bool struct {
	Pointer *bool
	Default bool
}

// NewBool returns a Bool struct as value which holds a pointer
// and a default value.
func NewBool(ptr *bool, def bool) Bool {
	return Bool{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to a bool using strconv.ParseBool
// and returns an error if that fails. Otherwise it sets the value
// of the pointer to the result of the conversion.
func (b Bool) Parse(str string) error {
	conv, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	*b.Pointer = conv
	return nil
}

// SetToDefault sets the value of the pointer to the default.
func (b Bool) SetToDefault() {
	*b.Pointer = b.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (b Bool) Value() interface{} {
	if b.Pointer == nil {
		return nil
	}

	return *b.Pointer
}
