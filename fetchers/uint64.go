package fetchers

import (
	"os"
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

func (i Uint64) Fetch(key string) error {
	var err error

	v := i.Default
	str, ok := os.LookupEnv(key)
	if ok {
		conv, parseErr := strconv.ParseUint(str, 10, 64)
		if parseErr == nil {
			v = conv
		} else {
			err = parseErr
		}
	} else {
		err = ErrNotSet
	}

	*i.Pointer = v
	return err
}

func (i Uint64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
