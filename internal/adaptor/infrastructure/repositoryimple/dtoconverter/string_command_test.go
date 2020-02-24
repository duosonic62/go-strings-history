package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/volatiletech/null"
	"strconv"
	"testing"
)

func TestToStringModel(t *testing.T) {
	gauge, _ := valueobject.NewStringGauge(9, 42)
	guitarString, err := entity.NewGuitarString("dummy_id", "dummy_name", "dummy_description", "dummy_maker", gauge, "")
	if err != nil {
		t.Error("Error occurred: ", err)
	}
	actual := ToStringModel(guitarString, 1)
	if actual.Name != "dummy_name" {
		t.Error("Expected: dummy_name, but " + actual.Name)
	}
	if !actual.Description.Valid || actual.Description.String != "dummy_description" {
		t.Error("Expected: dummy_description, but " + actual.Description.String)
	}
	if !actual.Maker.Valid || actual.Maker.String != "dummy_maker" {
		t.Error("Expected: dummy_maker, but " + actual.Maker.String)
	}
	if !actual.ThinGauge.Valid || actual.ThinGauge.Uint8 != uint8(9) {
		t.Error("Expected: 9, but " + string(actual.ThinGauge.Uint8))
	}
	if !actual.ThickGauge.Valid || actual.ThickGauge.Uint8 != 42 {
		t.Error("Expected: 42, but " + string(actual.ThickGauge.Uint8))
	}
	if actual.URL.String != "" {
		t.Error("Expected: url is null")
	}
	if actual.IsDeleted {
		t.Error("Expected: IsDeleted is false")
	}
	if actual.Version != 1 {
		t.Error("Expected: version is 1")
	}
}

func TestToStringEntity_Positive(t *testing.T) {
	stringModel := models.GuitarString{
		ID:          "dummy_id",
		Name:        "dummy_name",
		Description: null.NewString("dummy_description", true),
		Maker:       null.NewString("dummy_maker", true),
		ThinGauge:   null.NewUint8(uint8(10), true),
		ThickGauge:  null.NewUint8(uint8(46), true),
		URL:         null.NewString("dummy_url", true),
	}

	actual, err := ToStringEntity(&stringModel)
	if err != nil {
		t.Error("Error occurred: ", err)
	}

	if actual.GetName() != "dummy_name" {
		t.Error("Expected: dummy_name, but " + actual.GetName())
	}

	if actual.GetDescription() != "dummy_description" {
		t.Error("Expected: dummy_description, but " + actual.GetDescription())
	}

	if actual.GetMaker() != "dummy_maker" {
		t.Error("Expected: dummy_maker, but " + actual.GetMaker())
	}

	if actual.GetUrl() != "dummy_url" {
		t.Error("Expected: dummy_url, but " + actual.GetUrl())
	}

	if actual.GetGauge().ThinGauge() != 10 {
		t.Error("Expected: 10, but " + strconv.Itoa(actual.GetGauge().ThinGauge()))
	}

	if actual.GetGauge().ThickGauge() != 46 {
		t.Error("Expected: 46, but " + strconv.Itoa(actual.GetGauge().ThickGauge()))
	}
}

func TestToStringEntity_Negative(t *testing.T) {
	stringModel := models.GuitarString{
		ID:          "dummy_id",
		Name:        "dummy_name",
		Description: null.NewString("dummy_description", true),
		Maker:       null.NewString("dummy_maker", true),
		ThinGauge:   null.NewUint8(uint8(46), true),
		ThickGauge:  null.NewUint8(uint8(10), true),
		URL:         null.NewString("dummy_url", true),
	}

	_, err := ToStringEntity(&stringModel)
	if err == nil {
		t.Error("Error not occurred: ", err)
	}
}