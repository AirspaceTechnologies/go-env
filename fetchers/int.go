package fetchers

type Int struct {
	Pointer *int
	Default int
}

func NewInt(ptr *int, def int) Int {
	return Int{
		Pointer: ptr,
		Default: def,
	}
}

func (i Int) Fetch(key string) error {
	var i64 int64
	err := NewInt64(&i64, int64(i.Default)).Fetch(key)

	*i.Pointer = int(i64)

	return err
}

func (i Int) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
