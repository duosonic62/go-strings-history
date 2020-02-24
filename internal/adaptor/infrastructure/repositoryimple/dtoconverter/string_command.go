package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/volatiletech/null"
)

func ToStringModel(guitarString *entity.GuitarString, version int) *models.GuitarString {
	return &models.GuitarString{
		ID:          guitarString.GetID(),
		Name:        guitarString.GetName(),
		Description: null.StringFrom(guitarString.GetDescription()),
		Maker:       null.StringFrom(guitarString.GetMaker()),
		ThinGauge:   null.Uint8From(uint8(guitarString.GetGauge().ThinGauge())),
		ThickGauge:  null.Uint8From(uint8(guitarString.GetGauge().ThickGauge())),
		URL:         null.StringFrom(guitarString.GetUrl()),
		IsDeleted:   false,
		Version:     version,
	}
}

func ToStringEntity(stringModel *models.GuitarString) (*entity.GuitarString, error) {
	gauge, err := valueobject.NewStringGauge(int(stringModel.ThinGauge.Uint8), int(stringModel.ThickGauge.Uint8))
	if err != nil {
		return nil, err
	}

	entity, err := entity.NewGuitarString(
		stringModel.ID,
		stringModel.Name,
		stringModel.Description.String,
		stringModel.Maker.String,
		gauge,
		stringModel.URL.String,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
