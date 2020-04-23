package env

var (
	DefaultLogger LogFunc
	SuccessLogger LogFunc
	FailureLogger LogFunc
)

type LogFunc = func(string, ...interface{})

// use ZeroLogger to make sure something does not get logged
func ZeroLogger(string, ...interface{}) {}

// Config
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
