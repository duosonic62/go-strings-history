package inputboundary

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type UserQueryUseCase interface {
	Show(token valueobject.AuthorizationToken)
}
