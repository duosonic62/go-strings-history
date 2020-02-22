package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/null"
	"testing"
	"time"
)

func TestToStringOutput(t *testing.T) {
	guitarStringModel := models.GuitarString{
		ID: "mock_id",
		Name: "mock_name",
		Description: null.NewString("mock_description", true),
		Maker: null.NewString("mock_maker", true),
		ThinGauge: null.NewUint8(10, true),
		ThickGauge: null.NewUint8(46, true),
		URL: null.NewString("mock_url", true),
		IsDeleted: false,
		Version: 1,
		CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC),
		UpdatedAt: null.NewTime(time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC), true),
	}

	actual := ToStringOutput(&guitarStringModel)
	expected := output.GuitarStringOutput{
		ID: "mock_id",
		Name: "mock_name",
		Description: "mock_description",
		Maker: "mock_maker",
		ThinGauge: 10,
		ThickGauge: 46,
		Url: "mock_url",
	}

	if actual.ID != expected.ID {
		t.Error("ID expected " , expected.ID , ", bad ", actual.ID)
	}
	if actual.Maker != expected.Maker {
		t.Error("Maker expected " , expected.Maker , ", bad ", actual.Maker)
	}
	if actual.Description != expected.Description {
		t.Error("Description expected " , expected.Description , ", bad ", actual.Description)
	}
	if actual.ThinGauge != expected.ThinGauge {
		t.Error("ThinGauge expected " , expected.ThinGauge , ", bad ", actual.ThinGauge)
	}
	if actual.ThickGauge != expected.ThickGauge {
		t.Error("ThickGauge expected " , expected.ThickGauge , ", bad ", actual.ThickGauge)
	}
	if actual.Url != expected.Url {
		t.Error("Url expected " , expected.Url , ", bad ", actual.Url)
	}
}

func TestToStringOutputs(t *testing.T) {
	guitarString := models.GuitarString{
		ID: "mock_id",
		Name: "mock_name",
		Description: null.NewString("mock_description", true),
		Maker: null.NewString("mock_maker", true),
		ThinGauge: null.NewUint8(10, true),
		ThickGauge: null.NewUint8(46, true),
		URL: null.NewString("mock_url", true),
		IsDeleted: false,
		Version: 1,
		CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC),
		UpdatedAt: null.NewTime(time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC), true),
	}

	var guitarStringModels models.GuitarStringSlice
	guitarStringModels = make([]*models.GuitarString, 2)
	guitarStringModels[0] = &guitarString
	guitarStringModels[1] = &guitarString

	actuals := ToStringOutputs(&guitarStringModels)
	expected := output.GuitarStringOutput{
		ID: "mock_id",
		Name: "mock_name",
		Description: "mock_description",
		Maker: "mock_maker",
		ThinGauge: 10,
		ThickGauge: 46,
		Url: "mock_url",
	}

	for _, actual := range *actuals {
		if actual.ID != expected.ID {
			t.Error("ID expected " , expected.ID , ", bad ", actual.ID)
		}
		if actual.Maker != expected.Maker {
			t.Error("Maker expected " , expected.Maker , ", bad ", actual.Maker)
		}
		if actual.Description != expected.Description {
			t.Error("Description expected " , expected.Description , ", bad ", actual.Description)
		}
		if actual.ThinGauge != expected.ThinGauge {
			t.Error("ThinGauge expected " , expected.ThinGauge , ", bad ", actual.ThinGauge)
		}
		if actual.ThickGauge != expected.ThickGauge {
			t.Error("ThickGauge expected " , expected.ThickGauge , ", bad ", actual.ThickGauge)
		}
		if actual.Url != expected.Url {
			t.Error("Url expected " , expected.Url , ", bad ", actual.Url)
		}
	}
}