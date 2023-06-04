package rbconfig

type ConfigInterface interface {
	GetName() string
	Validate() error
}
