package parsers

import (
	"strconv"
)

// Int64 is the env.Parser implementation for the int64 type.
type Int64 struct {
	Pointer *int64
	Default int64
}

// NewInt64 returns an Int64 struct as value which holds a pointer
// and a default value.
func NewInt64(ptr *int64, def int64) Int64 {
	return Int64{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to an int64 using strconv.ParseInt
// and returns an error if that fails. Otherwise it sets the value
// of the pointer to the result of the conversion.
func (i Int64) Parse(str string) error {
	conv, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*i.Pointer = conv
	return err
}

// SetToDefault sets the value of the pointer to the default.
func (i Int64) SetToDefault() {
	*i.Pointer = i.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (i Int64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
