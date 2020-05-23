package env

var (
	DefaultLogger LogFunc
	SuccessLogger LogFunc
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

func DefaultConfig() Config {
	return Config{
		DefaultLogger: DefaultLogger,
		SuccessLogger: SuccessLogger,
		FailureLogger: FailureLogger,
	}
}
