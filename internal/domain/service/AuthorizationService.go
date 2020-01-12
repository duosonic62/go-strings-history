package service

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type AuthorizationService interface {
	Authorized(token valueobject.AuthorizationToken) (entity.User, error)
}