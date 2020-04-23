package env

func Fetch(vars []Var) {
	FetchWithConfig(vars, DefaultConfig())
}

func FetchWithConfig(vars []Var, cfg Config) {
	for _, v := range vars {
		fetch(v, cfg)
	}
}

func FetchMap(vars map[string]Var) {
	FetchMapWithConfig(vars, DefaultConfig())
}

func FetchMapWithConfig(vars map[string]Var, cfg Config) {
	for k, v := range vars {
		fetch(v.WithKey(k), cfg)
	}
}

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
