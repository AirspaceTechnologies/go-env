package parsers

// String is the env.Parser implementation for the string type.
type String struct {
	Pointer *string
	Default string
}

// NewString returns a String struct as value which holds a pointer
// and a default value.
func NewString(ptr *string, def string) String {
	return String{
		Pointer: ptr,
		Default: def,
	}
}

// Parse sets the value of the pointer to the string. It does not
// fail or convert since the value is already a string. The error
// returned will always be nil.
func (s String) Parse(str string) error {
	*s.Pointer = str
	return nil
}

// SetToDefault sets the value of the pointer to the default.
func (s String) SetToDefault() {
	*s.Pointer = s.Default
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (s String) Value() interface{} {
	if s.Pointer == nil {
		return nil
	}

	return *s.Pointer
}
