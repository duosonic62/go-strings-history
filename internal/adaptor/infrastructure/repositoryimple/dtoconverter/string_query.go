package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

func ToStringOutput(from *models.GuitarString) *output.GuitarStringOutput {
	output := &output.GuitarStringOutput{
		ID: from.ID,
		Name: from.Name,
		Description: from.Description.String,
		Maker: from.Maker.String,
		ThinGauge: int(from.ThinGauge.Uint8),
		ThickGauge: int(from.ThickGauge.Uint8),
		Url: from.URL.String,
	}

	if from.Description.Valid {
		output.Description = from.Description.String
	}

	if from.Maker.Valid {
		output.Maker = from.Maker.String
	}

	if from.ThinGauge.Valid && from.ThickGauge.Valid {
		output.ThinGauge = int(from.ThinGauge.Uint8)
		output.ThickGauge = int(from.ThickGauge.Uint8)
	}

	if from.URL.Valid {
		output.Url = from.URL.String
	}

	return output
}

func ToStringOutputs(froms *models.GuitarStringSlice) *[]output.GuitarStringOutput {
	outputs := make([]output.GuitarStringOutput, len(*froms))
	for _, from := range *froms {
		outputs = append(outputs, *ToStringOutput(from))
	}

	return &outputs
}