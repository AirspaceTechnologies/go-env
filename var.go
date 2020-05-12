package env

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/airspacetechnologies/go-env/parsers"
)

type Var struct {
	Key           string
	Parser        Parser
	DefaultLogger LogFunc
	SuccessLogger LogFunc
	FailureLogger LogFunc
	SetRequired   bool // true makes failure logger get called if env is not set
	Sensitive     bool // true prints **** instead of real value
}

func NewVar(k string, p Parser) Var {
	return Var{
		Key:    k,
		Parser: p,
	}
}

func BoolVar(k string, ptr *bool, def bool) Var {
	return NewVar(k, parsers.NewBool(ptr, def))
}

func DurationVar(k string, ptr *time.Duration, def time.Duration) Var {
	return NewVar(k, parsers.NewDuration(ptr, def))
}

func Float64Var(k string, ptr *float64, def float64) Var {
	return NewVar(k, parsers.NewFloat64(ptr, def))
}

func IntVar(k string, ptr *int, def int) Var {
	return NewVar(k, parsers.NewInt(ptr, def))
}

func Int64Var(k string, ptr *int64, def int64) Var {
	return NewVar(k, parsers.NewInt64(ptr, def))
}

func Uint64Var(k string, ptr *uint64, def uint64) Var {
	return NewVar(k, parsers.NewUint64(ptr, def))
}

func StringVar(k string, ptr *string, def string) Var {
	return NewVar(k, parsers.NewString(ptr, def))
}

func (v Var) WithKey(k string) Var {
	v.Key = k
	return v
}

func (v Var) WithParser(p Parser) Var {
	v.Parser = p
	return v
}

func (v Var) MakeSensitive() Var {
	v.Sensitive = true
	return v
}

func (v Var) LogNotSetAsFailure() Var {
	v.SetRequired = true
	return v
}

func (v Var) WithDefaultLogger(f LogFunc) Var {
	v.DefaultLogger = f
	return v
}

func (v Var) WithSuccessLogger(f LogFunc) Var {
	v.SuccessLogger = f
	return v
}

func (v Var) WithFailureLogger(f LogFunc) Var {
	v.FailureLogger = f
	return v
}

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

func (v Var) success(err error) bool {
	if err == nil {
		return true
	}

	return errors.Is(err, ErrNotSet) && !v.SetRequired
}
