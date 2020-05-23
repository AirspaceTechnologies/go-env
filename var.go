package env

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/airspacetechnologies/go-env/parsers"
)

// Var is the struct that coordinates the fetching,
// parsing, and logging of environment variables.
type Var struct {
	Key           string // the key to the environment variable
	Parser        Parser // the parser in charge of parsing and setting the variable
	DefaultLogger LogFunc
	SuccessLogger LogFunc
	FailureLogger LogFunc
	SetRequired   bool // true makes failure logger get called if env is not set
	Sensitive     bool // true prints **** instead of real value
}

// NewVar returns a Var struct with the key and parser set.
func NewVar(k string, p Parser) Var {
	return Var{
		Key:    k,
		Parser: p,
	}
}

// BoolVar is a convenience method to set up a new Var with a bool parser.
func BoolVar(k string, ptr *bool, def bool) Var {
	return NewVar(k, parsers.NewBool(ptr, def))
}

// DurationVar is a convenience method to set up a new Var with a time.Duration parser.
func DurationVar(k string, ptr *time.Duration, def time.Duration) Var {
	return NewVar(k, parsers.NewDuration(ptr, def))
}

// Float64Var is a convenience method to set up a new Var with a float64 parser.
func Float64Var(k string, ptr *float64, def float64) Var {
	return NewVar(k, parsers.NewFloat64(ptr, def))
}

// IntVar is a convenience method to set up a new Var with an int parser.
func IntVar(k string, ptr *int, def int) Var {
	return NewVar(k, parsers.NewInt(ptr, def))
}

// Int64Var is a convenience method to set up a new Var with an int64 parser.
func Int64Var(k string, ptr *int64, def int64) Var {
	return NewVar(k, parsers.NewInt64(ptr, def))
}

// Uint64Var is a convenience method to set up a new Var with a uint64 parser.
func Uint64Var(k string, ptr *uint64, def uint64) Var {
	return NewVar(k, parsers.NewUint64(ptr, def))
}

// StringVar is a convenience method to set up a new Var with a string parser.
func StringVar(k string, ptr *string, def string) Var {
	return NewVar(k, parsers.NewString(ptr, def))
}

// WithKey is a chainable method that returns a copy of the Var
// with Key set.
func (v Var) WithKey(k string) Var {
	v.Key = k
	return v
}

// WithParser is a chainable method that returns a copy of the Var
// with Parser set.
func (v Var) WithParser(p Parser) Var {
	v.Parser = p
	return v
}

// MakeSensitive is a chainable method that returns a copy of the Var
// with Sensitive set to true.
func (v Var) MakeSensitive() Var {
	v.Sensitive = true
	return v
}

// LogNotSetAsFailure is a chainable method that returns a copy of the Var
// with SetRequired set to true.
func (v Var) LogNotSetAsFailure() Var {
	v.SetRequired = true
	return v
}

// WithDefaultLogger is a chainable method that returns a copy of the Var
// with DefaultLogger set.
func (v Var) WithDefaultLogger(f LogFunc) Var {
	v.DefaultLogger = f
	return v
}

// WithSuccessLogger is a chainable method that returns a copy of the Var
// with SuccessLogger set.
func (v Var) WithSuccessLogger(f LogFunc) Var {
	v.SuccessLogger = f
	return v
}

// WithFailureLogger is a chainable method that returns a copy of the Var
// with FailureLogger set.
func (v Var) WithFailureLogger(f LogFunc) Var {
	v.FailureLogger = f
	return v
}

// Fetch looks up, parsers, sets, and logs variable from the environment.
//
// It first looks up the variable from the environment and passes the resulting
// string to the parser's Parse method if present. If the variable was not set
// in the environment or there was an error returned by the parser, the parser's
// SetToDefault method is called. Finally, the variable gets logged with the
// appropriate logging func according to the error and configuration. See the
// log function of Var for more details on logging.
func (v Var) Fetch() {
	var err error

	str, ok := os.LookupEnv(v.Key)
	if ok {
		err = v.Parser.Parse(str)
	} else {
		err = ErrNotSet
	}

	if err != nil {
		v.Parser.SetToDefault()
	}

	v.log(err)
}

// log logs the Var according to the configuration.
// It retrieves the appropriate logger and returns
// if there is none. Otherwise it creates the format
// string and passes that along with the rest of the
// arguments to the logging function.
//
// If the Var has Sensitive set to true then the
// actual variable will not be logged.
//
// If an error is present then it is included in the
// log message.
//
// See the logger function of Var for more details
// on the logging function that gets used.
func (v Var) log(err error) {
	f := v.logger(err)
	if f == nil {
		return
	}

	format := "set %v=%v"
	args := []interface{}{v.Key}

	if v.Sensitive {
		args = append(args, "****, value is filtered")
	} else {
		args = append(args, v.Parser.Value())
	}

	if err != nil {
		addition := "err: %v"
		if errors.Is(err, ErrNotSet) {
			addition = "default was used - %v"
		}

		format = fmt.Sprintf("%v, %v", format, addition)
		args = append(args, err)
	}

	f(format, args...)
}

// logger returns the correct logger depending on if the error
// is considered a success. The DefaultLogger is returned if
// the error is considered a success and there is no SuccessLogger
// or it's considered a failure and there is no failure logger.
//
// See the success function of Var for more details on what is
// considered a success.
func (v Var) logger(err error) func(string, ...interface{}) {
	var f func(string, ...interface{})

	if v.success(err) {
		f = v.SuccessLogger
	} else {
		f = v.FailureLogger
	}

	if f == nil {
		f = v.DefaultLogger
	}

	return f
}

// success returns a bool of whether or not the error is
// considered successful.
//
// The conditions that are considered successful are if
// the error is nil or the error is ErrNotSet and the Var
// has SetRequired set to false.
func (v Var) success(err error) bool {
	if err == nil {
		return true
	}

	return errors.Is(err, ErrNotSet) && !v.SetRequired
}
