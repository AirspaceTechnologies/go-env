package parsers

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

func (i Int) Parse(str string) error {
	var i64 int64
	if err := NewInt64(&i64, int64(i.Default)).Parse(str); err != nil {
		return err
	}

	*i.Pointer = int(i64)
	return nil
}

func (i Int) SetToDefault() {
	*i.Pointer = i.Default
}

func (i Int) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
