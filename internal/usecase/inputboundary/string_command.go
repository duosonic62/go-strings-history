package inputboundary

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type StringCommandUseCase interface {
	Add(data command.StringRegisterInputData, token *valueobject.AuthorizationToken, ctx input.Context)
	Update(id string, data command.StringRegisterInputData, token *valueobject.AuthorizationToken, ctx input.Context)
}