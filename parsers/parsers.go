// parsers is a package containing type specific logic used by the env package via the
// env.Parser interface. Parsers are responsible for converting a string to the correct
// type, setting the value of its pointer for both the parsed value and default value,
// and is able to return the value of its pointer as an interface or nil if the pointer
// is nil.
//
// Parsers for custom types can be created by following the implementation of the types
// in this package. The following is a demonstration of how to create a parser for a
// custom type (you can also view any of the parsers or the Parser example in the env package
// for further examples). I use generics as a stand in for a specific type in the example
// (hopefully generics come to golang soon!).
//  // create a struct to hold a pointer and default value
//	type CustomParser<T> struct {
//		Pointer *T
//		Default T
//	}
//
//	// not required but nice for convenience
//	func NewCustomParser<T>(ptr *T, def T) CustomParser<T> {
//		return CustomParser{
//			Pointer: ptr,
//			Default: def,
//		}
//	}
//
//  // Parse converts the string and sets the pointer upon success.
//  // If it fails it returns an error.
//	func (p CustomParser<T>) Parse(str string) error {
// 		// define or use a function to convert the string to your type
//		conv, err := someFuncThatConvertsTypeFromString(str)
//		if err != nil {
//			return err
//		}
//
//		*p.Pointer = conv
//		return nil
//	}
//
//  // SetToDefault gets called if the environmental variable was
//	// not set or if Parse returned an error. It just sets the
//	// value of the pointer.
//	func (p CustomParser<T>) SetToDefault() {
//		*p.Pointer = p.Default
//	}
//
// 	Value returns the value of the pointer or nil as an
// 	interface{} and is used for logging.
//	func (p CustomParser<T>) Value() interface{} {
//		if p.Pointer == nil {
//			return nil
//		}
//
//		return *p.Pointer
//	}
//
// You can also wrap a parser to add validation. See the ValidationParser
// example in the env package for how to easily do that.
package parsers
