package env

// Fetch is a convenience method for FetchWithConfig if you
// want to use the default config.
func Fetch(vars []Var) {
	FetchWithConfig(vars, DefaultConfig())
}

// FetchWithConfig iterates over a Var slice fetching each with the config.
func FetchWithConfig(vars []Var, cfg Config) {
	for _, v := range vars {
		fetch(v, cfg)
	}
}

// FetchMap is a convenience method for FetchMapWithConfig if you
// want to use the default config.
func FetchMap(vars map[string]Var) {
	FetchMapWithConfig(vars, DefaultConfig())
}

// FetchWithConfig iterates over a Var map
// fetching each key from the environment with the config.
func FetchMapWithConfig(vars map[string]Var, cfg Config) {
	for k, v := range vars {
		fetch(v.WithKey(k), cfg)
	}
}

// fetch sets the loggers on var and calls Fetch on Var.
func fetch(v Var, cfg Config) {
	if v.DefaultLogger == nil {
		v.DefaultLogger = cfg.DefaultLogger
	}

	if v.SuccessLogger == nil {
		v.SuccessLogger = cfg.SuccessLogger
	}

	if v.FailureLogger == nil {
		v.FailureLogger = cfg.FailureLogger
	}

	v.Fetch()
}
