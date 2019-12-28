package factoryimple

import (
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/google/uuid"
)

type TokenFactoryImpl struct {
}

func NewTokenFactory() factory.TokenFactory {
	return TokenFactoryImpl{}
}

func (factory TokenFactoryImpl) Gen() (string, error)  {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}