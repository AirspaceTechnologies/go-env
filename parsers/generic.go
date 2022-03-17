package parsers

import "github.com/airspacetechnologies/go-env/validators"

type Generic[T any] struct {
	Pointer      *T
	Default      T
	ParseFunc    ParseFunc[T]
	ValidateFunc validators.Func[T]
}

func NewGeneric[T any](ptr *T, def T, p ParseFunc[T], vfs ...validators.Func[T]) Generic[T] {
	parser := Generic[T]{
		Pointer:   ptr,
		Default:   def,
		ParseFunc: p,
	}
	return parser.AddValidators(vfs...)
}

// Parse converts the string to a type using Generic func
// and returns an error if that fails. Otherwise, it sets
// the value of the pointer to the result of the conversion.
func (p Generic[T]) Parse(str string) error {
	conv, err := p.ParseFunc(str)
	if err != nil {
		return err
	}

	if err := p.Validate(conv); err != nil {
		return err
	}

	*p.Pointer = conv
	return nil
}

// SetToDefault sets the value of the pointer to the default.
func (p Generic[T]) SetToDefault() {
	*p.Pointer = p.Default
}

// Validate runs all the validator functions that were added
// until an error is returned from one. if there are none
// then it will return nil.
func (p Generic[T]) Validate(v T) error {
	if p.ValidateFunc == nil {
		return nil
	}

	return p.ValidateFunc(v)
}

func (p Generic[T]) AddValidators(vfs ...validators.Func[T]) Generic[T] {
	for _, vf := range vfs {
		p.ValidateFunc = p.ValidateFunc.Wrap(vf)
	}

	return p
}

// Value returns the value of the pointer or nil as
// an interface{}.
func (p Generic[T]) Value() interface{} {
	if p.Pointer == nil {
		return nil
	}

	return *p.Pointer
}

// TypedValue returns the value of the pointer if the pointer is
// not nil. If the pointer is nil, it will return the default value.
func (p Generic[T]) TypedValue() T {
	if p.Pointer == nil {
		return p.Default
	}

	return *p.Pointer
}
