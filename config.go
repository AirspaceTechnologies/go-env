package env

var (
	// DefaultLogger logger is used for success if SuccessLogger
	// is not specified or for failures when FailureLogger
	// is not specified.
	DefaultLogger LogFunc

	// SuccessLogger is used when env var parsing was a success
	SuccessLogger LogFunc

	// FailureLogger is used when env var parsing was a failure
	FailureLogger LogFunc
)

// LogFunc is a type alias for a logging function
// such as log.Printf
type LogFunc = func(string, ...interface{})

// ZeroLogger can be used to make sure something does not get logged.
func ZeroLogger(string, ...interface{}) {}

// Config holds the logging functions.
type Config struct {
	DefaultLogger LogFunc
	SuccessLogger LogFunc
	FailureLogger LogFunc
}

// DefaultConfig is the config used if none other was specified
func DefaultConfig() Config {
	return Config{
		DefaultLogger: DefaultLogger,
		SuccessLogger: SuccessLogger,
		FailureLogger: FailureLogger,
	}
}
