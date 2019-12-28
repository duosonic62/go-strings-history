package factory

type TokenFactory interface {
	Gen() (string, error)
}