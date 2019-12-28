package factory

type IDFactory interface {
	Gen() (string, error)
}