package fetchers

import (
	"os"
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

func (i Int64) Fetch(key string) error {
	var err error

	v := i.Default
	str, ok := os.LookupEnv(key)
	if ok {
		conv, parseErr := strconv.ParseInt(str, 10, 64)
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

func (i Int64) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
