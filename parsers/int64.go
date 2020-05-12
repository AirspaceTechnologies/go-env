package parsers

import (
	"strconv"
)

type Int64 struct {
	Pointer *int64
	Default int64
}

func NewInt64(ptr *int64, def int64) Int64 {
	return Int64{
		Pointer: ptr,
		Default: def,
	}
}

func (i Int64) Parse(str string) error {
	conv, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*i.Pointer = conv
	return err
}

func (i Int64) SetToDefault() {
	*i.Pointer = i.Default
}

func (i Int64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
