package factory

import "github.com/duosonic/go-strings-history/internal/domain/entity"

type UserFactory struct {
	idFactory    IDFactory
	tokenFactory TokenFactory
}

// コンストラクタ
func NewUserFactory(idFactory IDFactory, tokenFactory TokenFactory) UserFactory {
	return UserFactory{
		idFactory:    idFactory,
		tokenFactory: tokenFactory,
	}
}

func (factory UserFactory) NewUser(name string, uid string) (entity.User, error) {
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
