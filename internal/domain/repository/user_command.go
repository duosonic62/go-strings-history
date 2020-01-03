package repository

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type UserCommandRepository interface {
	Find(valueobject.AuthorizationToken) (entity.User, error)
	Save(entity.User) error
}
