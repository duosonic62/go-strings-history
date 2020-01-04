package repository

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserQueryRepository interface {
	Find(valueobject.AuthorizationToken) (output.UserOutput, error)
}