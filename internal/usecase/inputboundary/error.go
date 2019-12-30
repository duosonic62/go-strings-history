package inputboundary

import "github.com/duosonic62/go-strings-history/pkg/usecase/input"

type ErrorUseCase interface {
	BadRequestError(ctx input.Context, err error)
	UnauthorizedError(ctx input.Context, err error)
}