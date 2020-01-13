package factoryimple

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

type StringFactoryImpl struct {
}

func (factory StringFactoryImpl) NewString(
	user *entity.User,
	thinGauge int,
	thickGauge int,
	name, description, maker, url string,
) (*entity.GuitarString, error) {
	gauge, err := valueobject.NewStringGauge(thinGauge, thickGauge)
	if err != nil {
		return nil, err
	}

	return entity.NewGuitarString(name, description, maker, gauge, url, user)
}
