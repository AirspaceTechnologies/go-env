package fetchers

import (
	"os"
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

func (d Duration) Fetch(key string) error {
	var err error

	v := d.Default
	str, ok := os.LookupEnv(key)
	if ok {
		conv, parseErr := time.ParseDuration(str)
		if parseErr == nil {
			v = conv
		} else {
			err = parseErr
		}
	} else {
		err = ErrNotSet
	}

	*d.Pointer = v
	return err
}

func (d Duration) Value() interface{} {
	if d.Pointer == nil {
		return nil
	}

	return *d.Pointer
}
