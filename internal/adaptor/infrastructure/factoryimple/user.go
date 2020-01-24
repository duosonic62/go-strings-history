package factoryimple

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type UserFactoryImpl struct {
	idFactory    factory.IDFactory
	tokenFactory factory.TokenFactory
	repository repository.UserCommandRepository
}

// コンストラクタ
func NewUserFactory(idFactory factory.IDFactory, tokenFactory factory.TokenFactory, repository repository.UserCommandRepository) factory.UserFactory {
	return UserFactoryImpl{
		idFactory:    idFactory,
		tokenFactory: tokenFactory,
		repository: repository,
	}
}

func (factory UserFactoryImpl) NewUser(name string, uid string) (*entity.User, error) {
	// idを生成
	id, err := factory.idFactory.Gen()
	if err != nil {
		return nil, entity.NewApplicationError(500, "id generate error", "Internal Server Error", err)
	}
	//tokenを生成
	tokenSeed, err := factory.tokenFactory.Gen()
	if err != nil {
		return nil, entity.NewApplicationError(500, "token generate error", "Internal Server Error", err)
	}

	token, err := valueobject.NewAuthorizationToken(tokenSeed)
	if err != nil {
		return nil, entity.NewApplicationError(500, "token generate error", "Internal Server Error", err)
	}

	user, err := entity.NewUser(id, name, uid, token)
	if err != nil {
		return nil, entity.NewApplicationError(500, "token generate error", "Internal Server Error", err)
	}

	return user, nil
}

func (factory UserFactoryImpl) Find(token *valueobject.AuthorizationToken) (*entity.User, error) {
	user, err := factory.repository.Find(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}