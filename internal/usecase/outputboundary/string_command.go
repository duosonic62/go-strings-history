package outputboundary

import "github.com/duosonic62/go-strings-history/pkg/usecase/input"

type StringCommandPresenter interface {
	OutputAddString(ctx input.Context)
}