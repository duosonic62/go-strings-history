package inputboundary

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type UserCommandUseCase interface {
	Add(data command.UserAddInputData, ctx input.Context)
	Edit(token *valueobject.AuthorizationToken, data command.UserEditInputData, ctx input.Context)
	Delete(token *valueobject.AuthorizationToken, ctx input.Context)
}