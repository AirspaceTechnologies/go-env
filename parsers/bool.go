package parsers

import (
	"strconv"
)

type Bool struct {
	Pointer *bool
	Default bool
}

func NewBool(ptr *bool, def bool) Bool {
	return Bool{
		Pointer: ptr,
		Default: def,
	}
}

func (b Bool) Parse(str string) error {
	conv, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	*b.Pointer = conv
	return nil
}

func (b Bool) SetToDefault() {
	*b.Pointer = b.Default
}

func (b Bool) Value() interface{} {
	if b.Pointer == nil {
		return nil
	}

	return *b.Pointer
}
