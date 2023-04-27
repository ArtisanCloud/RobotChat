package rbtemplate

type Template interface {
	GetName() string
	Execute(input string) string
}
