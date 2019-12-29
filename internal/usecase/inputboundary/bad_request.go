package inputboundary

import "github.com/duosonic62/go-strings-history/pkg/usecase/input"

type BadRequestUseCase interface {
	BadRequestError(ctx input.Context, err error)
}