package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
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