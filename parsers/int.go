package parsers

// Int is the env.Parser implementation for the int type.
type Int struct {
	Pointer *int
	Default int
}

// NewInt returns an Int struct as value which holds a pointer
// and a default value.
func NewInt(ptr *int, def int) Int {
	return Int{
		Pointer: ptr,
		Default: def,
	}
}

// Parse converts the string to an int using Int64 parser then converting
// the resulting int64 to an int returning an error if that fails.
// Otherwise it sets the value of the pointer to the result of the conversion.
func (i Int) Parse(str string) error {
	var i64 int64
	if err := NewInt64(&i64, int64(i.Default)).Parse(str); err != nil {
		return err
	}

	*i.Pointer = int(i64)
	return nil
}

// SetToDefault sets the value of the pointer to the default.
func (i Int) SetToDefault() {
	*i.Pointer = i.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (i Int) Value() interface{} {
	if i.Pointer == nil {
		return nil
	}

	return *i.Pointer
}
