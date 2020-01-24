package factory

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type UserFactory interface {
	NewUser(name string, uid string) (*entity.User, error)
	Find(token *valueobject.AuthorizationToken) (*entity.User, error)
}
