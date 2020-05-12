package env

type Parser interface {
	Parse(string) error
	SetToDefault()
	Value() interface{}
}
