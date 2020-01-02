package outputboundary

import (
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserQueryPresenter interface {
	OutputUser(output output.UserOutput,ctx input.Context)
}