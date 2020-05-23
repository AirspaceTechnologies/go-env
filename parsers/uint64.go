package parsers

import (
	"strconv"
)

// Uint64 is the env.Parser implementation for the unit64 type.
type Uint64 struct {
	Pointer *uint64
	Default uint64
}

// NewUint64 returns a Uint64 struct as value which holds a pointer
// and a default value.
func NewUint64(ptr *uint64, def uint64) Uint64 {
	return Uint64{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to a uint64 using strconv.ParseUint
// and returns an error if that fails. Otherwise it sets the value
// of the pointer to the result of the conversion.
func (i Uint64) Parse(str string) error {
	conv, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	*i.Pointer = conv
	return err
}

// SetToDefault sets the value of the pointer to the default.
func (i Uint64) SetToDefault() {
	*i.Pointer = i.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (i Uint64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
