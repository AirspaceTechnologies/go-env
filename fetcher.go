package env

type Fetcher interface {
	Fetch(key string) error
	Value() interface{}
}
