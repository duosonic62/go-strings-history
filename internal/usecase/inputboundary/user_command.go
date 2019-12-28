package inputboundary

import (
	"github.com/duosonic/go-strings-history/pkg/usecase/input"
	"github.com/duosonic/go-strings-history/pkg/usecase/input/command"
)

type UserCommandUseCase interface {
	AddUser(data command.UserAddInputData, ctx input.Context)
}