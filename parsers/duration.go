package parsers

import (
	"time"
)

// Duration is the env.Parser implementation for the time.Duration type.
type Duration struct {
	Pointer *time.Duration
	Default time.Duration
}

// NewDuration returns a Duration struct as value which holds a pointer
// and a default value.
func NewDuration(ptr *time.Duration, def time.Duration) Duration {
	return Duration{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to a time.Duration using time.ParseDuration
// and returns an error if that fails. Otherwise it sets the value of
// the pointer to the result of the conversion.
func (d Duration) Parse(str string) error {
	conv, err := time.ParseDuration(str)
	if err != nil {
		return err
	}

	*d.Pointer = conv
	return err
}

// SetToDefault sets the value of the pointer to the default.
func (d Duration) SetToDefault() {
	*d.Pointer = d.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (d Duration) Value() interface{} {
	if d.Pointer == nil {
		return nil
	}

	return *d.Pointer
}
