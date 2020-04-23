package fetchers

import (
	"os"
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

func (b Bool) Fetch(key string) error {
	var err error

	v := b.Default
	str, ok := os.LookupEnv(key)
	if ok {
		conv, parseErr := strconv.ParseBool(str)
		if parseErr == nil {
			v = conv
		} else {
			err = parseErr
		}
	} else {
		err = ErrNotSet
	}

	*b.Pointer = v
	return err
}

func (b Bool) Value() interface{} {
	if b.Pointer == nil {
		return nil
	}

	return *b.Pointer
}
