package rbconfig

type Config interface {
	GetName() string
	Validate() error
}
