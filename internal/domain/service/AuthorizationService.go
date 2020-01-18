package service

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type AuthorizationService interface {
	Authorized(token *valueobject.AuthorizationToken) (*entity.User, error)
}

type AuthorizationServiceImpl struct {
	userFactory factory.UserFactory
}

func NewAuthorizationService(userFactory factory.UserFactory) AuthorizationService {
	return AuthorizationServiceImpl{
		userFactory: userFactory,
	}
}

func (service AuthorizationServiceImpl) Authorized(token *valueobject.AuthorizationToken) (*entity.User, error) {
	user, err := service.userFactory.Find(token)
	if err != nil {
		return nil, entity.NewUnAuthorizedError(token, err)
	}

	return user, nil
}
