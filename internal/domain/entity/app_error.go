package entity

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type ApplicationError struct {
	Code  int
	Debug string
	Front string
	Cause error
}

func NewApplicationError(code int, debug, front string, cause error) ApplicationError {
	return ApplicationError{
		Code:  code,
		Debug: debug,
		Front: front,
		Cause: cause,
	}
}

func NewUnAuthorizedError(token *valueobject.AuthorizationToken, err error) ApplicationError {
	return ApplicationError{
		Code:  401,
		Debug: "this token is invalid: " + token.GetToken(),
		Front: "this token is invalid",
		Cause: err,
	}
}

func NewBadRequestError(err error) ApplicationError {
	return ApplicationError{
		Code:  400,
		Debug: "bad request" + err.Error(),
		Front: "bad request",
		Cause: err,
	}
}

func NewInternalServerError(err error) ApplicationError {
	return ApplicationError{
		Code:  500,
		Debug: "internal server error: " + err.Error(),
		Front: "internal server error",
		Cause: err,
	}
}

func (e ApplicationError) Error() string {
	return e.Debug
}
