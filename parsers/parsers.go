// parsers is a package containing type specific logic used by the env package via the
// env.Parser interface. Parsers are responsible for converting a string to the correct
// type, setting the value of its pointer for both the parsed value and default value,
// and is able to return the value of its pointer as an interface or nil if the pointer
// is nil.
//
// Parsers for custom types can be created by following the implementation of the types
// in this package. The following is a demonstration of how to create a parser for a
// custom type, I use generics as a stand in for a specific type in the example
// (hopefully generics come to golang soon!).
// 	type CustomParser<T> struct {
//		Pointer *T
//		Default T
//	}
//
//	func NewCustomParser<T>(ptr *T, def T) CustomParser<T> {
//		return CustomParser{
//			Pointer: ptr,
//			Default: def,
//		}
//	}
//
//	func (p CustomParser<T>) Parse(str string) error {
//		conv, err := someFuncThatConvertsTypeFromString(str)
//		if err != nil {
//			return err
//		}
//
//		*p.Pointer = conv
//		return nil
//	}
//
//	func (p CustomParser<T>) SetToDefault() {
//		*p.Pointer = p.Default
//	}
//
//	func (p CustomParser<T>) Value() interface{} {
//		if p.Pointer == nil {
//			return nil
//		}
//
//		return *p.Pointer
//	}
//
// You can also wrap a parser to add validation. The following is an example of how to do that.
// 	package main
//
//	import (
//		"errors"
//		"github.com/airspacetechnologies/go-env/parsers"
//	)
//
//	type PercentParser struct {
//		parsers.Float64
//	}
//
//	func (p PercentParser) Parse(str string) error {
//		if err := p.Float64.Parse(str); err != nil {
//			return err
//		}
//
//		v := *p.Pointer
//		if v < 0 || v > 100 {
//			return errors.New("percent is out of bounds")
//		}
//
//		return nil
//	}
package parsers
