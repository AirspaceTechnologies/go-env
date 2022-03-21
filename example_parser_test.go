package env_test

import (
	"encoding/hex"
	"fmt"
	"github.com/airspacetechnologies/go-env/validators"
	"net"
	"net/url"
	"os"

	"github.com/airspacetechnologies/go-env"
	"github.com/airspacetechnologies/go-env/parsers"
)

// NewHexParser is not required, it is just for convenience.
func NewHexParser(ptr *[]byte, def []byte) env.Parser {
	return parsers.NewGeneric(ptr, def, hex.DecodeString)
}

// ExampleHexParser shows how to make a custom parser.
func Example_hexParser() {
	key := "Example_hexParser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: newStdoutLogger(),
	}
	var b []byte

	// env variable not set - uses default
	v.WithParser(NewHexParser(&b, []byte{1})).Fetch()

	// env variable set to invalid hex string - uses default
	os.Setenv(key, "gg")

	v.WithParser(NewHexParser(&b, []byte{2})).Fetch()

	// env variable set to a valid hex string
	os.Setenv(key, "ff")

	v.WithParser(NewHexParser(&b, []byte{3})).Fetch()

	// Output: set Example_hexParser=[1], default was used - variable was not explicitly set in env
	// set Example_hexParser=[2], default was used - error: encoding/hex: invalid byte: U+0067 'g'
	// set Example_hexParser=[255]
}

// NewURLParser is not required, it is just for convenience.
func NewURLParser(ptr **url.URL, def *url.URL, vfs ...validators.Func[*url.URL]) env.Parser {
	return parsers.NewGeneric(ptr, def, url.Parse, vfs...)
}

// ExampleURLParser shows how to make a custom parser.
func Example_urlParser() {
	key := "Example_urlParser"
	defer os.Unsetenv(key)

	v := env.Var{
		Key:           key,
		DefaultLogger: newStdoutLogger(),
	}
	var u *url.URL

	defaultURL := &url.URL{
		Scheme:   "http",
		Host:     net.JoinHostPort("localhost", "8080"),
		Path:     "/test/path",
		RawQuery: "parma=value",
	}

	// env variable not set - uses default
	v.WithParser(NewURLParser(&u, defaultURL)).Fetch()

	// env variable set to invalid url string failing custom validator - uses default
	os.Setenv(key, "1")

	v.WithParser(NewURLParser(&u, defaultURL, func(u *url.URL) error {
		if u.Host == "" {
			return fmt.Errorf("missing host")
		}

		return nil
	})).Fetch()

	// env variable set to invalid url string - uses default
	os.Setenv(key, "1:://")

	v.WithParser(NewURLParser(&u, defaultURL)).Fetch()

	// env variable set to a valid url string
	os.Setenv(key, "postgres://user:password@localhost:5432/db?sslmode=disable")

	v.WithParser(NewURLParser(&u, defaultURL)).Fetch()

	// Output: set Example_urlParser=http://localhost:8080/test/path?parma=value, default was used - variable was not explicitly set in env
	// set Example_urlParser=http://localhost:8080/test/path?parma=value, default was used - error: missing host
	// set Example_urlParser=http://localhost:8080/test/path?parma=value, default was used - error: parse "1:://": first path segment in URL cannot contain colon
	// set Example_urlParser=postgres://user:password@localhost:5432/db?sslmode=disable
}
