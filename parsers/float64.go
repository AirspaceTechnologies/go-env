package parsers

import (
	"strconv"
)

type Float64 struct {
	Pointer *float64
	Default float64
}

func NewFloat64(ptr *float64, def float64) Float64 {
	return Float64{
		Pointer: ptr,
		Default: def,
	}
}

func (f Float64) Parse(str string) error {
	conv, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*f.Pointer = conv
	return err
}

func (f Float64) SetToDefault() {
	*f.Pointer = f.Default
}

func (f Float64) Value() interface{} {
	if f.Pointer == nil {
		return nil
	}

	return *f.Pointer
}
