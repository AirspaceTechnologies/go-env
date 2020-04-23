package fetchers

import (
	"os"
)

type String struct {
	Pointer *string
	Default string
}

func NewString(ptr *string, def string) String {
	return String{
		Pointer: ptr,
		Default: def,
	}
}

func (s String) Fetch(key string) error {
	var err error

	v, ok := os.LookupEnv(key)
	if !ok {
		v = s.Default
		err = ErrNotSet
	}

	*s.Pointer = v
	return err
}

func (s String) Value() interface{} {
	if s.Pointer == nil {
		return nil
	}

	return *s.Pointer
}
