package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/volatiletech/null"
)

func ToStringModel(guitarString *entity.GuitarString, version int) *models.GuitarString {
	return &models.GuitarString{
		ID: guitarString.GetID(),
		Name: guitarString.GetName(),
		Description: null.StringFrom(guitarString.GetDescription()),
		Maker: null.StringFrom(guitarString.GetMaker()),
		ThinGauge: null.Uint8From(uint8(guitarString.GetGauge().ThinGauge())),
		ThickGauge: null.Uint8From(uint8(guitarString.GetGauge().ThickGauge())),
		URL: null.StringFrom(guitarString.GetUrl()),
		IsDeleted: false,
		Version: version,
	}
}