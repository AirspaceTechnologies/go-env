package fetchers

import "errors"

var (
	ErrNotSet = errors.New("variable was not explicitly set in env")
)
