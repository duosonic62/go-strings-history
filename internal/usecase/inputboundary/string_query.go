package inputboundary

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type StringQueryUseCase interface {
	GetGuitarString(stringID string, token *valueobject.AuthorizationToken, ctx input.Context)
}