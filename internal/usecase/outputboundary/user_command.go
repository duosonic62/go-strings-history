package outputboundary

import (
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserCommandPresenter interface {
	OutputAddUser(output output.UserAddOutputData, ctx input.Context)
	OutputEditUser(output output.UserEditOutputData, ctx input.Context)
	OutputDeleteUser(ctx input.Context)
}