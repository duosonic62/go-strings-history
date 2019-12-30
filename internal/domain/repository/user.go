package repository

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserRepository interface {
	Find(valueobject.AuthorizationToken) (output.UserOutput, error)
	Save(entity.User) error
}