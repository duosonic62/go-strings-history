package inputboundary

import "github.com/duosonic62/go-strings-history/pkg/usecase/input"

type ErrorUseCase interface {
	Error(ctx input.Context, err error)
}