package parsers

import (
	"strconv"
)

type Uint64 struct {
	Pointer *uint64
	Default uint64
}

func NewUint64(ptr *uint64, def uint64) Uint64 {
	return Uint64{
		Pointer: ptr,
		Default: def,
	}
}

func (i Uint64) Parse(str string) error {
	conv, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	*i.Pointer = conv
	return err
}

func (i Uint64) SetToDefault() {
	*i.Pointer = i.Default
}

func (i Uint64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
