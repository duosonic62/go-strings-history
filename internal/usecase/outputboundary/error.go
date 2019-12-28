package outputboundary

import "github.com/duosonic62/go-strings-history/pkg/usecase/input"

type ErrorPresenter interface {
	OutputError(ctx input.Context, err error)
}
