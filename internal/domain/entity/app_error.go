package entity

type ApplicationError struct {
	Code int
	Debug string
	Front string
	Cause error
}

func NewApplicationError(code int, debug, front string, cause error) ApplicationError {
	return ApplicationError{
		Code: code,
		Debug: debug,
		Front: front,
		Cause: cause,
	}
}

func (e ApplicationError) Error() string {
	return e.Debug
}
