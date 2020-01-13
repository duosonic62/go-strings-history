package factoryimple

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type StringFactoryImpl struct {
	idFactory factory.IDFactory
}

func NewStringFactory(idFactory factory.IDFactory) factory.StringFactory {
	return StringFactoryImpl{
		idFactory: idFactory,
	}
}

func (factory StringFactoryImpl) NewString(
	user *entity.User,
	thinGauge int,
	thickGauge int,
	name, description, maker, url string,
) (*entity.GuitarString, error) {
	id, err := factory.idFactory.Gen()
	if err != nil {
		return nil, err
	}

	gauge, err := valueobject.NewStringGauge(thinGauge, thickGauge)
	if err != nil {
		return nil, err
	}

	return entity.NewGuitarString(id, name, description, maker, gauge, url, user)
}
