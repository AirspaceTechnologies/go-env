package env

import "errors"

var (
	// ErrNotSet is the error that gets returned when os.Lookup does not return ok.
	ErrNotSet = errors.New("variable was not explicitly set in env")
)
