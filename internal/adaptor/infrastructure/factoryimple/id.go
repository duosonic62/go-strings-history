package factoryimple

import (
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/google/uuid"
)

type IDFactoryImpl struct {
}

func NewIdFactory() factory.IDFactory {
	return IDFactoryImpl{}
}

func (factory IDFactoryImpl) Gen() (string, error)  {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}