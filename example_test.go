package env_test

import (
	"github.com/airspacetechnologies/go-env"
	"log"
	"os"
)

func newStdoutLogger() env.LogFunc {
	return log.New(os.Stdout, "", 0).Printf // sends logs to os.Stdout for examples
}
