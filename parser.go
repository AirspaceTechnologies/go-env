package env

// Parser is the interface wraps the methods called
// by Var during the fetching process. It is responsible
// for parsing the string from the environment, setting
// the variable to the parsed out or default value, and
// returning the value.
//
// Parse gets called with a string representation of
// the variable if os.LookupEnv returns ok. This method
// should also set the value of its pointer to the
// variable to the result of parsing upon success.
// It returns an error if an error was encountered during
// parsing.
//
// SetToDefault gets called if os.LookupEnv does not
// return ok or Parse returned an error. This method
// should set tha value of its pointer to the variable
// to the default value.
//
// Value should return the value of its pointer to the
// variable or nil.
type Parser interface {
	Parse(string) error
	SetToDefault()
	Value() interface{}
}
