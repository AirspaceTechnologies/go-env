package parsers

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

func (s String) Parse(str string) error {
	*s.Pointer = str
	return nil
}

func (s String) SetToDefault() {
	*s.Pointer = s.Default
}

func (s String) Value() interface{} {
	if s.Pointer == nil {
		return nil
	}

	return *s.Pointer
}
