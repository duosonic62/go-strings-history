package factory

import "github.com/duosonic62/go-strings-history/internal/domain/entity"

type UserFactory interface {
	NewUser(name string, uid string) (entity.User, error)
}

type UserFactoryImpl struct {
	idFactory    IDFactory
	tokenFactory TokenFactory
}

// コンストラクタ
func NewUserFactory(idFactory IDFactory, tokenFactory TokenFactory) UserFactory {
	return UserFactoryImpl{
		idFactory:    idFactory,
		tokenFactory: tokenFactory,
	}
}

func (factory UserFactoryImpl) NewUser(name string, uid string) (entity.User, error) {
	// idを生成
	id, err := factory.idFactory.Gen()
	if err != nil {
		return entity.User{}, err
	}
	//tokenを生成
	token, err := factory.tokenFactory.Gen()
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		ID:    id,
		Name:  name,
		UID:   uid,
		Token: token,
	}

	return user, nil
}
