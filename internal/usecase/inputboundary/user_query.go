package inputboundary

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type UserQueryUseCase interface {
	Show(token valueobject.AuthorizationToken, ctx input.Context)
}
