package fetchers

import (
	"os"
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

func (f Float64) Fetch(key string) error {
	var err error

	v := f.Default
	str, ok := os.LookupEnv(key)
	if ok {
		conv, parseErr := strconv.ParseFloat(str, 64)
		if parseErr == nil {
			v = conv
		} else {
			err = parseErr
		}
	} else {
		err = ErrNotSet
	}

	*f.Pointer = v
	return err
}

func (f Float64) Value() interface{} {
	if f.Pointer == nil {
		return nil
	}

	return *f.Pointer
}
