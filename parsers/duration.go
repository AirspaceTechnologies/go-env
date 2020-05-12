package parsers

import (
	"time"
)

type Duration struct {
	Pointer *time.Duration
	Default time.Duration
}

func NewDuration(ptr *time.Duration, def time.Duration) Duration {
	return Duration{
		Pointer: ptr,
		Default: def,
	}
}

func (d Duration) Parse(str string) error {
	conv, err := time.ParseDuration(str)
	if err != nil {
		return err
	}

	*d.Pointer = conv
	return err
}

func (d Duration) SetToDefault() {
	*d.Pointer = d.Default
}

func (d Duration) Value() interface{} {
	if d.Pointer == nil {
		return nil
	}

	return *d.Pointer
}
